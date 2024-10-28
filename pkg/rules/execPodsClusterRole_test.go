package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_Pods_All_Star_Permissions(t *testing.T) {
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
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ExecPodsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_Incorrect_Pod_Exec_Permissions(t *testing.T) {
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
  - pod/exec
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ExecPodsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_Pods_Pods_Exec_Verbs_Permissions(t *testing.T) {
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
  - pods/exec
  verbs:
  - get
  - create
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ExecPodsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Pods_Exec_Verbs_Create_Only(t *testing.T) {
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
  - pods/exec
  verbs:
  - create
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ExecPodsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_Pods_Exec_Verbs_Get_Only(t *testing.T) {
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
  - pods/exec
  verbs:
  - get
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ExecPodsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_Pods_Exec_Verbs_Split_Rules(t *testing.T) {
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
  verbs:
  - get
- apiGroups:
  - ""
  resources:
  - pods/exec
  verbs:
  - create
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ExecPodsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Pods_Verbs_Get_Only(t *testing.T) {
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
  verbs:
  - get
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ExecPodsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_Pods_Pods_Exec_Star_Permissions(t *testing.T) {
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
  - pods/exec
  verbs:
  - "*"
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ExecPodsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}
