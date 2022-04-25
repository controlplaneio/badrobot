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

	securityContext := RunAsNonRoot(json)
	if securityContext != 3 {
		t.Errorf("Got %v securityContext wanted %v", securityContext, 3)
	}
}

func Test_RunAsNonRoot_Deploy_Spec(t *testing.T) {
	var data = `
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
        runAsNonRoot: true
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	securityContext := RunAsNonRoot(json)
	if securityContext != 0 {
		t.Errorf("Got %v securityContext wanted %v", securityContext, 0)
	}
}
