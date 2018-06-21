package main

import (
	"gopkg.in/yaml.v2"
)

type config struct {
	Definitions  definitions  `yaml:"definitions"`
	Applications applications `yaml:"applications"`
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
definitions: 
  - name  : Version
	find  : (^|[^\.\d])(\d+\.\d+\.\d+\.?\d*)([^\.\d]|$)
    group : 2		
applications: 
  - variable : Version
    filemask : .*
    find     : (^|[^\.\d])(\d+\.\d+\.\d+\.?\d*)([^\.\d]|$)
    replace  : ${1}%s${3}
`

	return json
}
