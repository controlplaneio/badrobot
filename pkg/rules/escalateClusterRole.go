// OPR-R16-RBAC - ClusterRole has escalate permissions
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func EscalateClusterRole(input []byte) (int, error) {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return -1, err
	}

	for _, rule := range clusterRole.Rules {
		if contains("rbac.authorization.k8s.io", rule.APIGroups) &&
			contains("clusterroles", rule.Resources) &&
			contains("escalate", rule.Verbs) {
			rbac++
		}
	}

	return rbac, nil

}
