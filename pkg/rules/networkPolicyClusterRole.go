// OPR-R25-RBAC - ClusterRole has read, write or delete permissions over network policies
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func NetworkPolicyClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("networking.k8s.io", rule.APIGroups) &&
			containsAny([]string{"networkpolicy", "networkpolicies", "*"}, rule.Resources) &&
			containsAny([]string{"*", "create", "update", "patch", "delete", "deletecollection"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
