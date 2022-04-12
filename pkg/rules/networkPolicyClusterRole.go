// OPR-R24-RBAC - ClusterRole has full permissions over network policies
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func NetworkPolicyClusterRole(json []byte) int {
	rbac := 0

	jqAPI := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("apiGroups")

	jqResources := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("resources")

	jqVerbs := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("verbs")

	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[networking.k8s.io]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "*")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[networking.k8s.io]]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "networkpolicy")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[networking.k8s.io]]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "networkpolicy")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[networking.k8s.io]]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "networkpolicy")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "update")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[networking.k8s.io]]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "networkpolicy")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "patch")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[networking.k8s.io]]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "networkpolicy")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "delete")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[networking.k8s.io]]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "networkpolicy")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "deletecollection")) {
		rbac++
	}

	return rbac

}
