package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_Events_All_Permissions(t *testing.T) {
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
  - events
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := RemoveEventsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Incorrect_Events_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - authorization.k8s.io
  resources:
  - events
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := RemoveEventsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_Events_Verbs_Permissions(t *testing.T) {
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
  - events
  verbs:
  - delete
  - deletecollection
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := RemoveEventsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Events_Multiple_Resources(t *testing.T) {
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
  - events
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := RemoveEventsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Events_Multiple_API_Groups(t *testing.T) {
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
  - events
  - deployments
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := RemoveEventsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Events_Multiple_Rules(t *testing.T) {
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
  - events
  verbs:
  - "*"
- apiGroups:
  - "apps"
  resources:
  - deployments
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := RemoveEventsClusterRole(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}
