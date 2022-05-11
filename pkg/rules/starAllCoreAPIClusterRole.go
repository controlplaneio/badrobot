// OPR-R12-RBAC - ClusterRole has full permissions over all CoreAPI resources
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func StarAllCoreAPIClusterRole(input []byte) (int, error) {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return -1, err
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			contains("*", rule.Resources) &&
			contains("*", rule.Verbs) {
			rbac++
		} else if contains("", rule.APIGroups) &&
			contains("*", rule.Resources) &&
			containsAll([]string{
				"get",
				"create",
				"update",
				"list",
				"patch",
				"watch",
				"delete",
				"deletecollection",
			}, rule.Verbs) {
			rbac++
		}
	}

	return rbac, nil
}
