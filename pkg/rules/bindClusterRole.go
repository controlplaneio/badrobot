// OPR-R17-RBAC - ClusterRole has bind permissions
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func BindClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("rbac.authorization.k8s.io", rule.APIGroups) &&
			contains("clusterroles", rule.Resources) &&
			contains("bind", rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
