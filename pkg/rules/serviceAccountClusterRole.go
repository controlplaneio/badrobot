// OPR-R22-RBAC - ClusterRole has permissions over service account token creation
package rules

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ServiceAccountClusterRole(json []byte) int {
	rbac := 0

	jqRules := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules")

	numElementsStr := fmt.Sprintf("%v", jqRules.Count())
	numElementsVar, _ := strconv.Atoi(numElementsStr)

	reSA := regexp.MustCompile(`(serviceaccounts):?[^/]`)
	reSAToken := regexp.MustCompile(`(serviceaccounts/token)`)

	for i := 1; i <= numElementsVar; i++ {
		apiGroups := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["apiGroups"])
		resources := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["resources"])
		verbs := fmt.Sprintf("%v", jqRules.Nth(i).(map[string]interface{})["verbs"])

		if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			reSA.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			reSA.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			reSAToken.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "*") {
			rbac++
		} else if strings.Contains(fmt.Sprintf("%v", apiGroups), "[]") &&
			reSAToken.MatchString(fmt.Sprintf("%v", resources)) &&
			strings.Contains(fmt.Sprintf("%v", verbs), "create") {
			rbac++
		}

	}

	return rbac

}

// 	rbac := 0

// 	jqAPI := gojsonq.New().Reader(bytes.NewReader(json)).
// 		From("rules").
// 		Only("apiGroups")

// 	jqResources := gojsonq.New().Reader(bytes.NewReader(json)).
// 		From("rules").
// 		Only("resources")

// 	jqVerbs := gojsonq.New().Reader(bytes.NewReader(json)).
// 		From("rules").
// 		Only("verbs")

// 	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
// 		(strings.Contains(fmt.Sprintf("%v", jqResources), "[serviceaccounts]")) &&
// 		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
// 		rbac++
// 	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
// 		(strings.Contains(fmt.Sprintf("%v", jqResources), "[serviceaccounts]")) &&
// 		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
// 		rbac++
// 	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
// 		(strings.Contains(fmt.Sprintf("%v", jqResources), "[serviceaccounts/token]")) &&
// 		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
// 		rbac++
// 	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
// 		(strings.Contains(fmt.Sprintf("%v", jqResources), "[serviceaccounts/token]")) &&
// 		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
// 		rbac++
// 	}

// 	return rbac

// }
