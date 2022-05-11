// OPR-R22-RBAC - ClusterRole has full permissions over admission controllers
package rules

import (
	"encoding/json"

	rbacv1 "k8s.io/api/rbac/v1"
)

func AdmissionControllerClusterRole(input []byte) (int, error) {
	rbac := 0

	clusterRole := &rbacv1.ClusterRole{}
	err := json.Unmarshal(input, clusterRole)
	if err != nil {
		return -1, err
	}

	for _, rule := range clusterRole.Rules {
		if contains("admissionregistration.k8s.io", rule.APIGroups) &&
			containsAny([]string{"mutatingwebhookconfigurations", "validatingwebhookconfigurations"}, rule.Resources) &&
			containsAny([]string{"*", "create", "patch", "update", "delete", "deletecollection"}, rule.Verbs) {
			rbac++
		}
	}

	return rbac, nil

}
