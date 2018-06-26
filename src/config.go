package main

import (
	"gopkg.in/yaml.v2"
)

type config struct {
	Variables variableDefs `yaml:"variables"`
	Tasks     taskDefs     `yaml:"tasks"`
}

func deserializeYAML(content []byte) (result config) {
	err := yaml.Unmarshal(content, &result)

	if err != nil {
		panic(err)
	}
	return
}

func loadConfig(file fileIo, filename string) config {
	json := file.readBytes(filename)
	return deserializeYAML(json)
}

func basicYAML() string {
	json := `
variables: 
  - name: Version
    find: ([^\.\d]|^)(\d+\.\d+\.\d+\.?\d*)([^\.\d]|$)
    group: 2		
tasks: 
  - variable: Version
    filemask: \.xyz
    find: (^|[^\.\d])(\d+\.\d+\.\d+\.?\d*)([^\.\d]|$)
    replace: ${1}%s${3}
`
	return json
}
