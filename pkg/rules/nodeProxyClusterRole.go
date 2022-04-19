// OPR-R26-RBAC - ClusterRole has permissions over the Kubernetes API server proxy
package rules

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func NodeProxyClusterRole(json []byte) int {
	rbac := 0

	jqRules := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules")

	numElementsStr := fmt.Sprintf("%v", jqRules.Count())
	numElementsVar, _ := strconv.Atoi(numElementsStr)

	reNodes := regexp.MustCompile(`(nodes):?[^/]`)
	reNodesProxy := regexp.MustCompile(`(nodes/proxy)`)

	for i := 1; i <= numElementsVar; i++ {
		apiGroups := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["apiGroups"])
		resources := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["resources"])
		verbs := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["verbs"])

		if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			reNodes.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			reNodes.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "get") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			reNodesProxy.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			reNodesProxy.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "get") &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") {
			rbac++
		}

	}

	return rbac

}
