// OPR-R16-RBAC - ClusterRole has escalate permissions
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func EscalateClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if containsAny([]string{"*", "rbac.authorization.k8s.io"}, rule.APIGroups) &&
			containsAny([]string{"*", "clusterroles"}, rule.Resources) &&
			containsAny([]string{"*", "escalate"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac

}
