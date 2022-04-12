// OPR-R11-RBAC - ClusterRole has full permissions over all CoreAPI resources
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func StarAllCoreAPIClusterRole(json []byte) int {
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

	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "*")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "*")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "get")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "update")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "list")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "patch")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "watch")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "delete")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "deletecollection")) {
		rbac++
	}

	return rbac

}
