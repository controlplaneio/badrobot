package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_Cluster_Admin_Permissions(t *testing.T) {
	var data = `
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manager-rolebinding
subjects:
- kind: ServiceAccount
  name: manager-rolebinding
  # Replace this with the namespace the operator is deployed in.
  namespace: system
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ClusterAdmin(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Incorrect_Cluster_Admin_Permissions(t *testing.T) {
	var data = `
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manager-rolebinding
subjects:
- kind: ServiceAccount
  name: manager-rolebinding
  # Replace this with the namespace the operator is deployed in.
  namespace: system
roleRef:
  kind: ClusterRole
  name: cluster-administrator
  apiGroup: rbac.authorization.k8s.io
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac, err := ClusterAdmin(json)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}
