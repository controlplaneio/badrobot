package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func DefaultNamespace(json []byte) int {
	namespace := 0

	jqNS := gojsonq.New().Reader(bytes.NewReader(json)).
		From("metadata.name").Get()

	jqDeploy := gojsonq.New().Reader(bytes.NewReader(json)).
		From("metadata.namespace").Get()

	jqCRB := gojsonq.New().Reader(bytes.NewReader(json)).
		From("subjects").
		Only("namespace")

	if jqNS != nil && strings.Contains(fmt.Sprintf("%v", jqNS), "kube-system") ||
		jqNS != nil && strings.Contains(fmt.Sprintf("%v", jqNS), "default") ||
		strings.Contains(fmt.Sprintf("%v", jqDeploy), "kube-system") ||
		strings.Contains(fmt.Sprintf("%v", jqDeploy), "default") ||
		strings.Contains(fmt.Sprintf("%v", jqCRB), "kube-system") ||
		strings.Contains(fmt.Sprintf("%v", jqCRB), "default") {
		namespace++
	}

	return namespace
}
