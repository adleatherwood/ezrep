package main

import (
	"regexp"
)

type variableDef struct {
	Name  string `yaml:"name"`
	Find  string `yaml:"find"`
	Group int    `yaml:"group"`
}

type variableDefs []variableDef

type variableExp struct {
	name  string
	find  *regexp.Regexp
	group int
}

type variableExps []variableExp

type variable struct {
	name  string
	value string
}

type variables []variable
type variableMap map[string]variable

func (vd variableDef) toExpression() variableExp {
	return variableExp{
		name:  vd.Name,
		find:  regexp.MustCompile(vd.Find),
		group: vd.Group,
	}
}

func (vds variableDefs) toExpressions() (result variableExps) {
	for _, vd := range vds {
		result = append(result, vd.toExpression())
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

func (ves variableExps) execute(contents []string) (result variables) {
	for _, c := range contents {
		for _, ve := range ves {
			b := []byte(c)
			if ve.find.Match(b) {
				value := ve.find.FindStringSubmatch(c)[ve.group]
				result = append(result, variable{ve.name, value})
			}
		}
	}
	return
}

func (vds variableDefs) execute(contents []string) variableMap {
	ves := vds.toExpressions()
	vs := ves.execute(contents)
	return vs.toMap()
}

func (vm variableMap) print(console consoleIo) {
	for _, v := range vm {
		console.writeLn("Found Variable -> %s = %s", v.name, v.value)
	}
}
