package rules

import (
	"github.com/ghodss/yaml"
	"testing"
)

func Test_Privileged_Pod(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: c1
    securityContext:
      privileged: true
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := Privileged(json)
	if containers != 1 {
		t.Errorf("Got %v containers wanted %v", containers, 1)
	}
}

func Test_Privileged_Missing(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: c1
    securityContext:
  - name: c2
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := Privileged(json)
	if containers != 0 {
		t.Errorf("Got %v containers wanted %v", containers, 0)
	}
}
