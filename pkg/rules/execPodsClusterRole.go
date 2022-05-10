// OPR-R15-RBAC - ClusterRole can exec into Pods
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func ExecPodsClusterRole(input []byte) int {
	rbac := 0

	var foundPodsGet, foundExecCreate bool

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			containsAll([]string{"pods", "pods/exec"}, rule.Resources) &&
			(contains("*", rule.Verbs) || containsAll([]string{"get", "create"}, rule.Verbs)) {
			rbac++
		} else if contains("", rule.APIGroups) &&
			contains("pods", rule.Resources) &&
			containsAny([]string{"*", "get"}, rule.Verbs) {
			foundPodsGet = true
			if foundPodsGet && foundExecCreate {
				rbac++
			}
		} else if contains("", rule.APIGroups) &&
			contains("pods/exec", rule.Resources) &&
			containsAny([]string{"*", "create"}, rule.Verbs) {
			foundExecCreate = true
			if foundPodsGet && foundExecCreate {
				rbac++
			}
		}

	}

	return rbac
}
