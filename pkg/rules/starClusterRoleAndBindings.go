// OPR-R13-RBAC - ClusterRole has full permissions over ClusterRoles and ClusterRoleBindings
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func StarClusterRoleAndBindings(input []byte) int {
	rbac := 0
	var foundCR, foundCRB bool

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("rbac.authorization.k8s.io", rule.APIGroups) &&
			containsAll([]string{"clusterroles", "clusterrolebindings"}, rule.Resources) &&
			(contains("*", rule.Verbs) || containsAll([]string{
				"get",
				"create",
				"update",
				"list",
				"patch",
				"watch",
				"delete",
				"deletecollection",
			}, rule.Verbs)) {
			rbac++
		} else if contains("rbac.authorization.k8s.io", rule.APIGroups) &&
			contains("clusterroles", rule.Resources) &&
			(contains("*", rule.Verbs) || containsAll([]string{
				"get",
				"create",
				"update",
				"list",
				"patch",
				"watch",
				"delete",
				"deletecollection",
			}, rule.Verbs)) {
			foundCR = true
			if foundCR && foundCRB {
				rbac++
			}
		} else if contains("rbac.authorization.k8s.io", rule.APIGroups) &&
			contains("clusterrolebindings", rule.Resources) &&
			(contains("*", rule.Verbs) || containsAll([]string{
				"get",
				"create",
				"update",
				"list",
				"patch",
				"watch",
				"delete",
				"deletecollection",
			}, rule.Verbs)) {
			foundCRB = true
			if foundCR && foundCRB {
				rbac++
			}
		}
	}

	return rbac
}
