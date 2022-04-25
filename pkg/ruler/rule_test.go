package ruler

import (
	"fmt"
	"testing"

	"github.com/controlplaneio/badrobot/pkg/rules"
	"github.com/ghodss/yaml"
)

func TestRule_Eval(t *testing.T) {
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
      - command:
        - /manager
        args:
        - --leader-elect
        image: controller:latest
        name: manager
        securityContext:
          allowPrivilegeEscalation: false
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rule := &Rule{
		Predicate: rules.AllowPrivilegeEscalation,
		Kinds:     []string{"Deployment"},
	}

	matchedSecurityContext, err := rule.Eval(json)
	if err != nil {
		t.Fatal(err.Error())
	}
	if matchedSecurityContext != 0 {
		fmt.Printf("%v", matchedSecurityContext)
		t.Errorf(fmt.Sprintf("Rule failed when it shouldn't with count %d", matchedSecurityContext))
	}
}

func TestRule_EvalDoesNotApply(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
kind: StatefulSet
spec:
  template:
    spec:
      containers:
        - name: manager
          image: controller:latest
          securityContext:
            allowPrivilegeEscalation: true
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rule := &Rule{
		Predicate: rules.AllowPrivilegeEscalation,
		Kinds:     []string{"Deployment"},
	}

	_, err = rule.Eval(json)
	if err == nil {
		t.Errorf("Rule succeeded when it shouldn't")
	}
}

func TestRule_EvalNoKind(t *testing.T) {
	var data = `
---
apiVersion: apps/v1
spec:
  template:
    spec:
      containers:
        - name: manager
          image: controller:latest
          securityContext:
            allowPrivilegeEscalation: true
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rule := &Rule{
		Predicate: rules.AllowPrivilegeEscalation,
		Kinds:     []string{"Deployment"},
	}

	_, err = rule.Eval(json)
	if err == nil {
		t.Errorf("Rule succeeded when it shouldn't")
	}
}
