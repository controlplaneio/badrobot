// OPR-R23-RBAC - ClusterRole has permissions over service account token creation
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func ServiceAccountClusterRole(input []byte) (int, error) {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return -1, err
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			contains("serviceaccounts/token", rule.Resources) &&
			containsAny([]string{"*", "create"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac, nil

}
