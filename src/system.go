package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type directoryFileLister func(directory string) []string

type directoryIo struct {
	listFiles directoryFileLister
}

type fileByteReader func(filename string) []byte
type fileByteWriter func(filename string, content []byte)

type fileIo struct {
	readBytes  fileByteReader
	writeBytes fileByteWriter
}

type consoleWriter func(message string, a ...interface{})
type consoleLineWriter func(message string, a ...interface{})

type consoleIo struct {
	write   consoleWriter
	writeLn consoleLineWriter
}

type systemIo struct {
	file      fileIo
	directory directoryIo
	console   consoleIo
}

func defaultSystemIo() systemIo {
	return systemIo{
		defaultFileIo(),
		defaultDirectoryIo(),
		defaultConsoleIo(),
	}
}

func previewSystemIo() systemIo {
	return systemIo{
		previewFileIo(),
		defaultDirectoryIo(),
		defaultConsoleIo(),
	}
}

func defaultDirectoryIo() directoryIo {
	return directoryIo{directoryListFiles}
}

func defaultConsoleIo() consoleIo {
	return consoleIo{consoleWrite, consoleWriteLn}
}

func defaultFileIo() fileIo {
	return fileIo{fileReadBytes, fileWriteBytes}
}

func previewFileIo() fileIo {
	return fileIo{fileReadBytes, func(f string, b []byte) {}}
}

func directoryListFiles(directory string) []string {
	files := make([]string, 0)
	err := filepath.Walk(directory, func(path string, file os.FileInfo, err error) error {
		files = append(files, path)
		return err
	})

	if err != nil {
		panic(err)
	}

	return files
}

func fileReadBytes(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return bytes
}

func fileWriteBytes(filename string, contents []byte) {
	err := ioutil.WriteFile(filename, contents, 0644)

	if err != nil {
		panic(err)
	}
}

func consoleWrite(message string, a ...interface{}) {
	fmt.Printf(message, a...)
}

func consoleWriteLn(message string, a ...interface{}) {
	fmt.Printf(message, a...)
	fmt.Println()
}
