package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_Namespace_No_Labels(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: operator-namespace
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	namespace := PodSecurityStandardProfile(json)
	if namespace != 0 {
		t.Errorf("Got %v namespaces wanted %v", namespace, 0)
	}
}

func Test_Namespace_PSS_Privileged(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: operator-namespace
  labels:
    pod-security.kubernetes.io/enforce: privileged
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	namespace := PodSecurityStandardProfile(json)
	if namespace != 0 {
		t.Errorf("Got %v namespaces wanted %v", namespace, 0)
	}
}

func Test_Namespace_PSS_Baseline(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: operator-namespace
  labels:
    pod-security.kubernetes.io/enforce: baseline
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	namespace := PodSecurityStandardProfile(json)
	if namespace != 1 {
		t.Errorf("Got %v namespaces wanted %v", namespace, 1)
	}
}

func Test_Namespace_PSS_Restricted(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: operator-namespace
  labels:
    pod-security.kubernetes.io/enforce: restricted
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	namespace := PodSecurityStandardProfile(json)
	if namespace != 1 {
		t.Errorf("Got %v namespaces wanted %v", namespace, 1)
	}
}

func Test_Namespace_PSS_Audit_Only(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: operator-namespace
  labels:
    pod-security.kubernetes.io/audit: baseline
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	namespace := PodSecurityStandardProfile(json)
	if namespace != 0 {
		t.Errorf("Got %v namespaces wanted %v", namespace, 0)
	}
}
