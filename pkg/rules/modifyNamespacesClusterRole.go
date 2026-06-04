// OPR-R27-RBAC - ClusterRole has modify permissions over namespaces
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func ModifyNamespacesClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if containsAny([]string{"", "*"}, rule.APIGroups) &&
			containsAny([]string{"namespaces", "*"}, rule.Resources) &&
			containsAny([]string{"patch", "update", "*"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
