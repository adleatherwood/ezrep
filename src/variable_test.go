package main

import (
	"testing"
)

func TestBasicExpression(t *testing.T) {
	d := definition{
		Name:  "",
		Find:  "\\d",
		Group: 0,
	}
	variableScenario(t, d, "->1<-", "1")
}

func TestExpressionWithGroups(t *testing.T) {
	d := definition{
		Name:  "",
		Find:  "(>)([a-z])(<)",
		Group: 2,
	}
	variableScenario(t, d, "->a<-", "a")
}

func variableScenario(t *testing.T, find definition, content string, expected string) {
	defs := definitions{find}
	contents := []string{content}
	variables := defs.execute(contents)
	actual := variables[find.Name].value

	if actual != expected {
		t.Logf("variableScenario Failed -> expected: %s, actual: %s, find: %s, content: %s", expected, actual, find.Find, content)
		t.Fail()
	}
}
