package rules

import (
	"github.com/ghodss/yaml"
	"testing"
)

func Test_ReadOnlyRootFilesystem(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      containers:
        - name: c1
        - name: c2
          securityContext:
            readOnlyRootFilesystem: false
        - name: c3
          securityContext:
            readOnlyRootFilesystem: true
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := ReadOnlyRootFilesystem(json)
	if containers != 1 {
		t.Errorf("Got %v containers wanted %v", containers, 1)
	}
}

func Test_ReadOnlyRootFilesystem_NotSpecified(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo
spec:
  containers:
  - name: sec-ctx-demo
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      capabilities:
        add:
          - CHOWN
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := ReadOnlyRootFilesystem(json)
	if containers != 0 {
		t.Errorf("Got %v containers wanted %v", containers, 0)
	}
}

func Test_ReadOnlyRootFilesystem_NoContainers(t *testing.T) {
	var data = `
---
apiVersion: extensions/v1beta1
kind: Deployment
spec:
  template:
    spec:
      serviceAccountName: badrobot
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := ReadOnlyRootFilesystem(json)
	if containers != 0 {
		t.Errorf("Got %v containers wanted %v", containers, 0)
	}
}
