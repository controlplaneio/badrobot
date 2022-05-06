// OPR-R14-RBAC - ClusterRole has access to Kubernetes secrets
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func SecretsClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			contains("secrets", rule.Resources) &&
			containsAny([]string{"*", "get", "create", "update", "list", "patch", "watch"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
