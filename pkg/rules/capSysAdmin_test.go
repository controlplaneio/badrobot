package rules

import (
	"github.com/ghodss/yaml"
	"testing"
)

func Test_CapSysAdmin_Pod(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: controller-manager
  namespace: system
  labels:
    control-plane: controller-manager
spec:
  selector:
    matchLabels:
      control-plane: controller-manager
  replicas: 1
  template:
    metadata:
      annotations:
        kubectl.kubernetes.io/default-container: manager
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - name: c1
        image: controller:latest
        securityContext:
          capabilities:
            add:
              - SYS_ADMIN
      - name: c2
        image: controller:latest
        securityContext:
          capabilities:
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := CapSysAdmin(json)
	if containers != 1 {
		t.Errorf("Got %v containers wanted %v", containers, 1)
	}
}

func Test_CapSysAdmin_Spec(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: controller-manager
  namespace: system
  labels:
    control-plane: controller-manager
spec:
  selector:
    matchLabels:
      control-plane: controller-manager
  replicas: 1
  template:
    metadata:
      annotations:
        kubectl.kubernetes.io/default-container: manager
      labels:
        control-plane: controller-manager
    spec:
      securityContext:
        capabilities:
          add:
          - SYS_ADMIN
      containers:
      - name: c1
        securityContext:
          capabilities:
            add:
              - SYS_ADMIN
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := CapSysAdmin(json)
	if containers != 2 {
		t.Errorf("Got %v containers wanted %v", containers, 2)
	}
}

func Test_CapSysAdmin_Missing(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Pod
spec:
  initContainers:
  - name: init1
  containers:
  - name: c1
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := CapSysAdmin(json)
	if containers != 0 {
		t.Errorf("Got %v containers wanted %v", containers, 0)
	}
}
