package rules

import (
	"github.com/ghodss/yaml"
	"testing"
)

func Test_RunAsNonRoot(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      containers:
        - name: c1
          securityContext:
            runAsNonRoot: false
        - name: c2
          securityContext:
            runAsNonRoot: false
        - name: c3
          securityContext:
            runAsNonRoot: false
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := RunAsNonRoot(json)
	if containers != 3 {
		t.Errorf("Got %v containers wanted %v", containers, 3)
	}
}
