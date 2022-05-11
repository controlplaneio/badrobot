package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_Escalate_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - clusterroles
  verbs:
  - escalate
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := EscalateClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Escalate_Permissions_Multiple_Rules(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - clusterroles
  verbs:
  - escalate
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := EscalateClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}
func Test_Incorrect_Escalate_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - serviceaccounts
  verbs:
  - escalate
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := EscalateClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}
