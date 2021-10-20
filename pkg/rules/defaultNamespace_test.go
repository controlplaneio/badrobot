package rules

import (
	"github.com/ghodss/yaml"
	"testing"
)

// -- TODO KGW Write tests

func Test_NonDefaultNamespace(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: example
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	operator := DefaultNamespace(json)
	if operator != 0 {
		t.Errorf("Got %v operator wanted %v", operator, 0)
	}
}

func Test_DefaultNamespace(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: kube-system
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	operator := DefaultNamespace(json)
	if operator != 1 {
		t.Errorf("Got %v operator wanted %v", operator, 1)
	}
}
