// OPR-R12-RBAC - ClusterRole has full permissions over ClusterRoles and ClusterRoleBindings
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func StarClusterRoleAndBindings(json []byte) int {
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

	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "rbac.authorization.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "clusterroles")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "clusterrolebindings")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	}

	return rbac

}
