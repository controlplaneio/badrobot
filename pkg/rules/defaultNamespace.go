// OPR-R1-NS - default namespace
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func DefaultNamespace(json []byte) (int, error) {
	namespace := 0

	jqNS := gojsonq.New().Reader(bytes.NewReader(json)).
		From("metadata.name").Get()

	jqDeploy := gojsonq.New().Reader(bytes.NewReader(json)).
		From("metadata.namespace").Get()

	jqCRB := gojsonq.New().Reader(bytes.NewReader(json)).
		From("subjects").
		Only("namespace")

	if strings.Contains(fmt.Sprintf("%v", jqNS), "default") ||
		strings.Contains(fmt.Sprintf("%v", jqDeploy), "default") ||
		strings.Contains(fmt.Sprintf("%v", jqCRB), "default") {
		namespace++
	}

	return namespace, nil
}
