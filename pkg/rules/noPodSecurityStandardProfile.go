package rules

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
)

func NoPodSecurityStandardProfile(input []byte) int {
	namespace := &corev1.Namespace{}
	err := json.Unmarshal(input, namespace)
	if err != nil {
		return 0
	}

	if pssEnforcLabelValue, ok := namespace.Labels["pod-security.kubernetes.io/enforce"]; ok {
		if pssEnforcLabelValue == "baseline" || pssEnforcLabelValue == "restricted" {
			return 0
		}
	}

	return 1
}
