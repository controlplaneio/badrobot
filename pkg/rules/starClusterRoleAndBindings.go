// OPR-R13-RBAC - ClusterRole has full permissions over ClusterRoles and ClusterRoleBindings
package rules

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func StarClusterRoleAndBindings(json []byte) int {
	rbac := 0

	jqRules := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules")

	numElementsStr := fmt.Sprintf("%v", jqRules.Count())
	numElementsVar, _ := strconv.Atoi(numElementsStr)

	for i := 1; i <= numElementsVar; i++ {
		apiGroups := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["apiGroups"])
		resources := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["resources"])
		verbs := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["verbs"])

		if strings.Contains(fmt.Sprintf("%v", apiGroups), "rbac.authorization.k8s.io") &&
			strings.Contains(fmt.Sprintf("%v", resources), "clusterroles") &&
			strings.Contains(fmt.Sprintf("%v", resources), "clusterrolebindings") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "rbac.authorization.k8s.io") &&
			strings.Contains(fmt.Sprintf("%v", resources), "clusterroles") &&
			strings.Contains(fmt.Sprintf("%v", resources), "clusterrolebindings") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "get") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "update") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "list") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "patch") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "watch") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "delete") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "deletecollection") {
			rbac++
		}

	}

	return rbac
}

// 	jqAPI := gojsonq.New().Reader(bytes.NewReader(json)).
// 		From("rules").
// 		Only("apiGroups")

// 	jqResources := gojsonq.New().Reader(bytes.NewReader(json)).
// 		From("rules").
// 		Only("resources")

// 	jqVerbs := gojsonq.New().Reader(bytes.NewReader(json)).
// 		From("rules").
// 		Only("verbs")

// 	if strings.Contains(fmt.Sprintf("%v", jqAPI), "rbac.authorization.k8s.io") &&
// 		strings.Contains(fmt.Sprintf("%v", jqResources), "clusterroles") &&
// 		strings.Contains(fmt.Sprintf("%v", jqResources), "clusterrolebindings") &&
// 		strings.Contains(fmt.Sprintf("%v", jqVerbs), "*") {
// 		rbac++
// 	}

// 	return rbac

// }
