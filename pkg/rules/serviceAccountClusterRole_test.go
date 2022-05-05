package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_ServiceAccount_All_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - ""
  resources:
  - serviceaccounts
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := ServiceAccountClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_ServiceAccount_Only_Get_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - ""
  resources:
  - serviceaccounts
  verbs:
  - get
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := ServiceAccountClusterRole(json)
	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_ServiceAccount_Token_Create_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - ""
  resources:
  - serviceaccounts/token
  verbs:
  - create
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := ServiceAccountClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_ServiceAccount_Multiple_API_Groups(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - ""
  - apps
  resources:
  - serviceaccounts
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := ServiceAccountClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_ServiceAccount_Multiple_Resources(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - serviceaccounts
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := ServiceAccountClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_ServiceAccount_Multiple_Rules(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - ""
  resources:
  - serviceaccounts
  verbs:
  - "*"
- apiGroups:
  - "apps"
  resources:
  - deployments
  verbs:
  - "*"
- apiGroups:
  - ""
  resources:
  - serviceaccounts/token
  verbs:
  - create
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := ServiceAccountClusterRole(json)
	if rbac != 2 {
		t.Errorf("Got %v permissions wanted %v", rbac, 2)
	}
}
