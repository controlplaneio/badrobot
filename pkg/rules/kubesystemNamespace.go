// OPR-R2-NS - KubeSystem Namespace

package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func KubeSystemNamespace(json []byte) int {
	namespace := 0

	jqNS := gojsonq.New().Reader(bytes.NewReader(json)).
		From("metadata.name").Get()

	jqDeploy := gojsonq.New().Reader(bytes.NewReader(json)).
		From("metadata.namespace").Get()

	jqCRB := gojsonq.New().Reader(bytes.NewReader(json)).
		From("subjects").
		Only("namespace")

	if strings.Contains(fmt.Sprintf("%v", jqNS), "kube-system") ||
		strings.Contains(fmt.Sprintf("%v", jqDeploy), "kube-system") ||
		strings.Contains(fmt.Sprintf("%v", jqCRB), "kube-system") {
		namespace++
	}

	return namespace
}
