package main

import (
	"regexp"
)

type definition struct {
	Name  string `yaml:"name"`
	Find  string `yaml:"find"`
	Group int    `yaml:"group"`
}

type definitions []definition

type expression struct {
	name  string
	find  *regexp.Regexp
	group int
}

type expressions []expression

type variable struct {
	name  string
	value string
}

type variables []variable
type variableMap map[string]variable

func (d definition) toExpression() expression {
	return expression{
		name:  d.Name,
		find:  regexp.MustCompile(d.Find),
		group: d.Group,
	}
}

func (ds definitions) toExpressions() (result expressions) {
	for _, d := range ds {
		result = append(result, d.toExpression())
	}
	return
}

func (vs variables) toMap() (result variableMap) {
	result = make(map[string]variable)
	for _, v := range vs {
		result[v.name] = v
	}
	return
}

func (es expressions) execute(contents []string) (result variables) {
	for _, c := range contents {
		for _, e := range es {
			b := []byte(c)
			if e.find.Match(b) {
				value := e.find.FindStringSubmatch(c)[e.group]
				result = append(result, variable{e.name, value})
			}
		}
	}
	return
}

func (ds definitions) execute(contents []string) variableMap {
	es := ds.toExpressions()
	vs := es.execute(contents)
	return vs.toMap()
}

func (vm variableMap) print(console consoleIo) {
	for _, v := range vm {
		console.writeLn("Found Variable -> %s", v.name, v.value)
	}
}
