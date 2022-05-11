// OPR-R18-RBAC - ClusterRole has impersonate permissions
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func ImpersonateClusterRole(input []byte) (int, error) {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return -1, err
	}

	for _, rule := range clusterRole.Rules {
		if containsAny([]string{"", "*"}, rule.APIGroups) &&
			contains("serviceaccounts", rule.Resources) &&
			contains("impersonate", rule.Verbs) {
			rbac++
		}
	}

	return rbac, nil
}
