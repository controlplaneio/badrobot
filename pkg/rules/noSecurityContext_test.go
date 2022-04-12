package rules

import (
	"github.com/ghodss/yaml"
	"testing"
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

	containers := ContainersSecurityContext(json)
	if containers != 3 {
		t.Errorf("Got %v containers wanted %v", containers, 3)
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

	containers := NoContainersSecurityContext(json)
	if containers != 3 {
		t.Errorf("Got %v containers wanted %v", containers, 3)
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

	containers := SpecSecurityContext(json)
	if containers != 3 {
		t.Errorf("Got %v containers wanted %v", containers, 3)
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

	containers := NoSpecSecurityContext(json)
	if containers != 2 {
		t.Errorf("Got %v containers wanted %v", containers, 2)
	}
}
