// OPR-R19-RBAC - ClusterRole can modify pod logs
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func ModifyPodLogsClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			contains("pods/log", rule.Resources) &&
			containsAny([]string{"*", "create", "patch", "update", "delete", "deletecollection"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
