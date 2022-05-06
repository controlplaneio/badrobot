// OPR-R24-RBAC - ClusterRole has read, write or delete permissions over persistent volumes
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func PersistentVolumeClusterRole(input []byte) int {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return 0
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			contains("persistentvolumeclaims", rule.Resources) &&
			containsAny([]string{"*", "get", "list", "create", "patch", "update", "delete", "deletecollection", "watch"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac
}
