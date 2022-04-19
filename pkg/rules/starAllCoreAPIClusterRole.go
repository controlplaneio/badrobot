// OPR-R12-RBAC - ClusterRole has full permissions over all CoreAPI resources
package rules

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func StarAllCoreAPIClusterRole(json []byte) int {
	rbac := 0

	jqRules := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules")

	numElementsStr := fmt.Sprintf("%v", jqRules.Count())
	numElementsVar, _ := strconv.Atoi(numElementsStr)

	for i := 1; i <= numElementsVar; i++ {
		apiGroups := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["apiGroups"])
		resources := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["resources"])
		verbs := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["verbs"])

		if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			strings.Contains(fmt.Sprintf("%v", resources), "*") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			strings.Contains(fmt.Sprintf("%v", resources), "*") &&
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
