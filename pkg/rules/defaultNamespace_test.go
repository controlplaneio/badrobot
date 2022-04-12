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
  name: system
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	namespace := NonDefaultNamespace(json)
	if namespace != 1 {
		t.Errorf("Got %v namespace wanted %v", namespace, 1)
	}
}

func Test_KubeSystemNamespace(t *testing.T) {
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

	namespace := KubeSystemNamespace(json)
	if namespace != 1 {
		t.Errorf("Got %v namespace wanted %v", namespace, 1)
	}
}

func Test_DefaultNamespace(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: default
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	namespace := DefaultNamespace(json)
	if namespace != 1 {
		t.Errorf("Got %v namespace wanted %v", namespace, 1)
	}
}
