// OPR-R26-RBAC - ClusterRole has permissions over the Kubernetes API server proxy
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func NodeProxyClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if containsAny([]string{"*", ""}, rule.APIGroups) &&
			containsAny([]string{"*", "nodes/proxy"}, rule.Resources) &&
			containsAny([]string{"*", "get"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
