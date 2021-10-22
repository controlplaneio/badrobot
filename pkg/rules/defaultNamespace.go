package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func DefaultNamespace(json []byte) int {
	operators := 0

	jqNS := gojsonq.New().Reader(bytes.NewReader(json)).
		From("metadata.name").Get()

	if jqNS != nil && strings.Contains(fmt.Sprintf("%v", jqNS), "kube-system") {
		operators++
	}

	return operators
}
