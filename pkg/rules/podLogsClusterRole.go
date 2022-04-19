// OPR-R18-RBAC - ClusterRole can modify pod logs
package rules

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ModifyPodLogsClusterRole(json []byte) int {
	rbac := 0

	jqRules := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules")

	numElementsStr := fmt.Sprintf("%v", jqRules.Count())
	numElementsVar, _ := strconv.Atoi(numElementsStr)

	rePods := regexp.MustCompile(`(pods):?[^/]`)
	rePodsLog := regexp.MustCompile(`(pods/log)`)

	for i := 1; i <= numElementsVar; i++ {
		apiGroups := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["apiGroups"])
		resources := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["resources"])
		verbs := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["verbs"])

		if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			rePods.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			rePods.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") ||
			strings.Contains(fmt.Sprintf("%v", verbs), "update") ||
			strings.Contains(fmt.Sprintf("%v", verbs), "patch") ||
			strings.Contains(fmt.Sprintf("%v", verbs), "delete") ||
			strings.Contains(fmt.Sprintf("%v", verbs), "deletecollection") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			rePodsLog.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			rePodsLog.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") ||
			strings.Contains(fmt.Sprintf("%v", verbs), "update") ||
			strings.Contains(fmt.Sprintf("%v", verbs), "patch") ||
			strings.Contains(fmt.Sprintf("%v", verbs), "delete") ||
			strings.Contains(fmt.Sprintf("%v", verbs), "deletecollection") {
			rbac++
		}

	}

	return rbac

}
