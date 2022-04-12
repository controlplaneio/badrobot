// OPR-R15-RBAC - ClusterRole has escalate permissions
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ModifyPodLogsClusterRole(json []byte) int {
	rbac := 0

	jqAPI := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("apiGroups")

	jqResources := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		// Select("resources").Get()
		Only("resources")

	jqVerbs := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("verbs")

	fmt.Sprintf("%v", jqResources)

	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "modify")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "delete")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "deletecollection")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "patch")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "update")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods/log]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods/log]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods/log]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "modify")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods/log]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "delete")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods/log]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "deletecollection")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods/log]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "patch")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[pods/log]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "update")) {
		rbac++
	}

	return rbac

}
