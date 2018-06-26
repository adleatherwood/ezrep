package main

import (
	"fmt"
	"regexp"
)

type taskDef struct {
	Variable string `yaml:"variable"`
	Filemask string `yaml:"filemask"`
	Find     string `yaml:"find"`
	Replace  string `yaml:"replace"`
}

type taskDefs []taskDef

type task struct {
	variable variable
	filemask *regexp.Regexp
	find     *regexp.Regexp
	replace  string
}

type tasks []task

type update struct {
	variable variable
	filename string
}

type updates []update

func (td taskDef) toTask(variable variable) task {
	return task{
		variable: variable,
		filemask: regexp.MustCompile(td.Filemask),
		find:     regexp.MustCompile(td.Find),
		replace:  td.Replace,
	}
}

func (tds taskDefs) toTasks(vars variableMap) (result tasks) {
	for _, td := range tds {
		variable := vars[td.Variable]
		result = append(result, td.toTask(variable))
	}
	return
}

func (t task) execute(file fileIo, filename string) update {
	content := file.readBytes(filename)
	replace := []byte(fmt.Sprintf(t.replace, t.variable.value))
	updated := t.find.ReplaceAll(content, replace)

	file.writeBytes(filename, updated)

	return update{t.variable, filename}
}

func (ts tasks) execute(file fileIo, filenames []string) (result updates) {
	for _, filename := range filenames {
		for _, task := range ts {
			b := []byte(filename)
			if task.filemask.Match(b) {
				update := task.execute(file, filename)
				result = append(result, update)
			}
		}
	}
	return
}

func (tds taskDefs) execute(file fileIo, variables variableMap, filenames []string) updates {
	tasks := tds.toTasks(variables)
	return tasks.execute(file, filenames)
}

func (u update) print(console consoleIo) {
	console.writeLn("File Matched -> variable: %s, file: %s", u.variable.name, u.filename)
}

func (us updates) print(console consoleIo) {
	for _, u := range us {
		u.print(console)
	}
}
