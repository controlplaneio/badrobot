package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_Nodes_All_Permissions(t *testing.T) {
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
  - nodes/proxy
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := NodeProxyClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Incorrect_Node_Permissions(t *testing.T) {
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
  - node
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := NodeProxyClusterRole(json)
	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_Nodes_Proxy_Verbs_Permissions(t *testing.T) {
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
  - nodes/proxy
  verbs:
  - get
  - create
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := NodeProxyClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Nodes_Proxy_Multiple_Rules(t *testing.T) {
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
  - nodes/proxy
  verbs:
  - get
  - create
- apiGroups:
  - apps
  resources:
  - deployments
  verbs:
  - "*"
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := NodeProxyClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Nodes_Proxy_Multiple_API_Groups(t *testing.T) {
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
  - nodes/proxy
  - deployments
  verbs:
  - get
  - create
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := NodeProxyClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}
