package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func DefaultNamespace(json []byte) int {
	operators := 0

	kubeSys := gojsonq.New().Reader(bytes.NewReader(json)).
		From("metadata.name").Get()

	if kubeSys != nil && strings.Contains(fmt.Sprintf("%v", kubeSys), "kube-system") {
		operators++
	}

	return operators
}
