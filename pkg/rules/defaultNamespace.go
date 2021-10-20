package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func DefaultNamespace(json []byte) int {
	spec := getSpecSelector(json)
	operators := 0

	kubeSys := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec + ".namespace").
		Only("metadata.name")

	if kubeSys != nil && strings.Contains(fmt.Sprintf("%v", kubeSys), "kube-system") {
		operators++
	}

	// capAddInit := gojsonq.New().Reader(bytes.NewReader(json)).
	// 	From(spec + ".initContainers").
	// 	Only("securityContext.capabilities.add")

	// if capAddInit != nil && strings.Contains(fmt.Sprintf("%v", capAddInit), "SYS_ADMIN") {
	// 	namespace++
	// }

	return operators
}
