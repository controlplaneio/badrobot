package rules

import (
	"github.com/ghodss/yaml"
	"testing"
)

func Test_RunAsUser_Zero_Pod(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: c1
    securityContext:
      runAsUser: 0
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	securityContext := RunAsUser(json)
	if securityContext != 1 {
		t.Errorf("Got %v securityContext wanted %v", securityContext, 1)
	}
}

func Test_RunAsUser_Zero_Deploy(t *testing.T) {
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
        runAsUser: 0
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	securityContext := RunAsUser(json)
	if securityContext != 1 {
		t.Errorf("Got %v securityContext wanted %v", securityContext, 1)
	}
}
