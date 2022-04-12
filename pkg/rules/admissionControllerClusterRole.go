// OPR-R21-RBAC - ClusterRole has full permissions over admission controllers
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func AdmissionControllerClusterRole(json []byte) int {
	rbac := 0

	jqAPI := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("apiGroups")

	jqResources := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("resources")

	jqVerbs := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("verbs")

	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "mutatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "mutatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "mutatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "patch")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "mutatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "update")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "mutatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "delete")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "mutatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "deletecollection")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "validatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "validatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "validatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "patch")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "validatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "update")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "validatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "delete")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[admissionregistration.k8s.io")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "validatingwebhookconfigurations")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "deletecollection")) {
		rbac++

	}

	return rbac

}
