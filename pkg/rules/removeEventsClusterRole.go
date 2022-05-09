// OPR-R20-RBAC - ClusterRole can remove Kubernetes events
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func RemoveEventsClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			contains("events", rule.Resources) &&
			containsAny([]string{"*", "delete", "deletecollection"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
