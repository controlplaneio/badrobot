// OPR-R26-RBAC - ClusterRole has permissions over the Kubernetes API server proxy
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func NodeProxyClusterRole(input []byte) (int, error) {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return -1, err
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			contains("nodes/proxy", rule.Resources) &&
			contains("*", rule.Verbs) {
			rbac++
		} else if contains("", rule.APIGroups) &&
			contains("nodes/proxy", rule.Resources) &&
			containsAll([]string{"get", "create"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac, nil
}
