package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_ContainersSecurityContext(t *testing.T) {
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
        allowPrivilegeEscalation: false
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	securityContext, err := NoSecurityContext(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if securityContext != 0 {
		t.Errorf("Got %v securityContext wanted %v", securityContext, 0)
	}
}

func Test_NoContainersSecurityContext(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      containers:
      - name: c1
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	securityContext, err := NoSecurityContext(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if securityContext != 1 {
		t.Errorf("Got %v securityContext wanted %v", securityContext, 1)
	}
}

func Test_SpecSecurityContext(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      securityContext:
        allowPrivilegeEscalation: false
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	securityContext, err := NoSecurityContext(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if securityContext != 0 {
		t.Errorf("Got %v securityContext wanted %v", securityContext, 0)
	}
}

func Test_NoSpecSecurityContext(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	securityContext, err := NoSecurityContext(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if securityContext != 1 {
		t.Errorf("Got %v securityContext wanted %v", securityContext, 1)
	}
}
