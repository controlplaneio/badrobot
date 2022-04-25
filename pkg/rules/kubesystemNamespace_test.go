package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

// -- TODO KGW Write tests

func Test_NonKubeSystemNamespace(t *testing.T) {
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

	namespace := KubeSystemNamespace(json)
	if namespace != 0 {
		t.Errorf("Got %v namespace wanted %v", namespace, 0)
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
		t.Errorf("Got %v namespace wanted %v", namespace, 0)
	}
}
