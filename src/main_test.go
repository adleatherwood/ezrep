package main

import (
	"testing"
)

func TestExampleScenario(t *testing.T) {
	config := `
variables: 
- name  : Version
  find  : ([vV]ersion=)([\d\.]+)(.*) 
  group : 2         
tasks: 
- variable : Version
  filemask : AssemblyInfo\.cs
  find     : ([vV]ersion\(")([\d\.]+)("\))
  replace  : ${1}%s${3}
- variable : Version
  filemask : \.[cf]sproj
  find     : (<[vV]ersion>)([\d\.]+)(<\/[vV]ersion>)
  replace  : ${1}%s${3}
  `

	files := make(FileMap)
	files[`path\Other.cs`] = fileResult{
		Original: "test",
		Expected: "test",
		Actual:   "test"}
	files[`path\AssemblyInfo.cs`] = fileResult{
		Original: `
		[assembly: AssemblyVersion("1.0.0.0")]
		[assembly: AssemblyFileVersion("1.0.0.0")]
		`,
		Expected: `
		[assembly: AssemblyVersion("1.2.3.4")]
		[assembly: AssemblyFileVersion("1.2.3.4")]
		`,
		Actual: ""}
	files[`path\MyProject.csproj`] = fileResult{
		Original: `
			...
			<Version>4.3.2.1</Version>
			...`,
		Expected: `
			...
			<Version>1.2.3.4</Version>
			...`,
		Actual: ""}

	inputs := []string{"Version=1.2.3.4"}

	doProcess(
		files.toDirectoryIo(),
		files.toFileIo(),
		toConfig(config),
		"<root=n/a>",
		inputs)

	for k, f := range files {
		if f.Expected != f.Actual {
			t.Logf("TestDoExecute Failed -> file: %s expected: %s, actual: %s", k, f.Expected, f.Actual)
			t.Fail()
		}
	}
}

type fileResult struct {
	Original string
	Expected string
	Actual   string
}

type FileMap map[string]fileResult

func toConfig(c string) config {
	return deserializeYAML([]byte(c))
}

func (files *FileMap) toDirectoryIo() directoryIo {
	return directoryIo{
		testLister(files),
	}
}

func (files *FileMap) toFileIo() fileIo {
	return fileIo{
		testReader(files),
		testwriter(files),
	}
}

func testLister(files *FileMap) directoryFileLister {
	return func(path string) []string {
		return files.toFilenames()
	}
}

func (files FileMap) toFilenames() (result []string) {
	for k := range files {
		result = append(result, k)
	}
	return
}

func testReader(files *FileMap) fileByteReader {
	return func(filename string) []byte {
		return []byte((*files)[filename].Original)
	}
}

func testwriter(files *FileMap) fileByteWriter {
	lambda := func(filename string, contents []byte) {
		current := (*files)[filename]
		(*files)[filename] = fileResult{current.Original, current.Expected, string(contents)}
	}
	return lambda
}
