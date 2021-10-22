package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func StarAllRoleRule(json []byte) int {
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
		(strings.Contains(fmt.Sprintf("%v", jqResources), "*")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	}

	fmt.Printf("%v", rbac)
	return rbac

}
