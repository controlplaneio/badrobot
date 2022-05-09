// OPR-R23-RBAC - ClusterRole has permissions over service account token creation
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func ServiceAccountClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			containsAny([]string{"serviceaccounts", "serviceaccounts/token"}, rule.Resources) &&
			containsAny([]string{"*", "create"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac

}
