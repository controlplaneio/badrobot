// OPR-R15-RBAC - ClusterRole can exec into Pods
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func ExecPodsClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			contains("pods/exec", rule.Resources) &&
			contains("*", rule.Verbs) {
			rbac++
		} else if contains("", rule.APIGroups) &&
			contains("pods/exec", rule.Resources) &&
			containsAny([]string{"create", "get"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
