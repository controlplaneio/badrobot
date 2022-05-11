// OPR-R21-RBAC - ClusterRole has full permissions over any custom resource definitions
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func CustomResourceClusterRole(input []byte) (int, error) {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return -1, err
	}

	for _, rule := range clusterRole.Rules {
		if contains("apiextensions.k8s.io", rule.APIGroups) &&
			contains("customresourcedefinitions", rule.Resources) &&
			containsAny([]string{"*", "create", "patch", "update", "delete", "deletecollection"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac, nil
}
