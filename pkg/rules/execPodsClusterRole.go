// OPR-R15-RBAC - ClusterRole can exec into Pods
package rules

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ExecPodsClusterRole(json []byte) int {
	rbac := 0

	jqRules := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules")

	numElementsStr := fmt.Sprintf("%v", jqRules.Count())
	numElementsVar, _ := strconv.Atoi(numElementsStr)

	rePods := regexp.MustCompile(`(pods):?[^/]`)
	rePodsExec := regexp.MustCompile(`(pods/exec)`)

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
			strings.Contains(fmt.Sprintf("%v", verbs), "get") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			rePodsExec.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			rePodsExec.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "get") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") {
			rbac++
		}

	}

	return rbac

}
