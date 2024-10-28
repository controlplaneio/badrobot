// OPR-R24-RBAC - ClusterRole has read, write or delete permissions over persistent volumes
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func PersistentVolumeClusterRole(input []byte) (int, error) {
	rbac := 0
	var foundPV, foundPVC bool

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return -1, err
	}

	for _, rule := range clusterRole.Rules {
		if contains("", rule.APIGroups) &&
			containsAll([]string{"persistentvolumes", "persistentvolumeclaims"}, rule.Resources) &&
			containsAny([]string{"*", "get", "list", "create", "patch", "update", "delete", "deletecollection", "watch"}, rule.Verbs) {
			rbac++
		} else if contains("", rule.APIGroups) &&
			contains("persistentvolumes", rule.Resources) &&
			containsAny([]string{"*", "get", "list", "create", "patch", "update", "delete", "deletecollection", "watch"}, rule.Verbs) {
			foundPV = true
			if foundPV && foundPVC {
				rbac++
			}
		} else if contains("", rule.APIGroups) &&
			contains("persistentvolumeclaims", rule.Resources) &&
			containsAny([]string{"*", "get", "list", "create", "patch", "update", "delete", "deletecollection", "watch"}, rule.Verbs) {
			foundPVC = true
			if foundPV && foundPVC {
				rbac++
			}
		}
	}

	return rbac, nil
}
