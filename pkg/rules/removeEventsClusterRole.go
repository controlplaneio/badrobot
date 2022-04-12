// OPR-R14-RBAC - ClusterRole can exec into Pods
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func RemoveEventsClusterRole(json []byte) int {
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

	// fmt.Printf("%v", jqAPI)
	// fmt.Printf("%v", jqResources)
	// fmt.Printf("%v", jqVerbs)

	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[events]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[events]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "delete")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "deletecollection")) {
		rbac++
	}

	return rbac

}
