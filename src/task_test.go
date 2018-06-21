package main

import (
	"testing"
)

type ezTask struct {
	variable string
	value    string
	find     string
	replace  string
}

func TestAssemblyInfoExpression(t *testing.T) {
	task := ezTask{
		variable: `Version`,
		value:    `1.2.3.4`,
		find:     `([vV]ersion\(")([\d\.]+)("\))`,
		replace:  `${1}%s${3}`,
	}.toTask()

	expected := `
	[assembly: AssemblyVersion("1.2.3.4")]
	[assembly: AssemblyFileVersion("1.2.3.4")]
	`
	content := `
	[assembly: AssemblyVersion("1.0.0.0")]
	[assembly: AssemblyFileVersion("1.0.0.0")]
	`
	taskScenario(t, task, expected, content)
}

func TestXsProjExpression(t *testing.T) {
	task := ezTask{
		variable: `Version`,
		value:    `1.2.3.4`,
		find:     `(<[vV]ersion>)([\d\.]+)(<\/[vV]ersion>)`,
		replace:  `${1}%s${3}`,
	}.toTask()

	expected := `<Version>1.2.3.4</Version>`
	content := `<Version>1.0.0.0</Version>`

	taskScenario(t, task, expected, content)
}

func taskScenario(t *testing.T, task task, expected string, content string) {
	file := testFileIo(content, func(f string, c []byte) {
		actual := string(c)
		if expected != actual {
			t.Logf("taskScenario Failed -> expected: %s, actual: %s", expected, actual)
			t.Fail()
		}
	})

	task.execute(file, "<test-file>")
}

func testFileIo(content string, writer fileByteWriter) fileIo {
	return fileIo{
		readBytes:  func(filename string) []byte { return []byte(content) },
		writeBytes: writer,
	}
}

func (t ezTask) toTask() task {
	application := application{
		Variable: t.variable,
		Filemask: "",
		Find:     t.find,
		Replace:  t.replace,
	}
	variable := variable{
		name:  t.variable,
		value: t.value,
	}

	return application.toTask(variable)
}
