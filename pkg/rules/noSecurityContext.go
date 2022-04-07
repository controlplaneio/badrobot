package rules

import (
	"bytes"

	"github.com/thedevsaddam/gojsonq/v2"
)

func NoSecurityContext(json []byte) int {
	spec := getSpecSelector(json)

	jqContainers := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".containers").
		Where("securityContext", "==", nil)

	jqSecurityContext := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec).
		Where("securityContext", "==", nil)

	return jqContainers.Count() //+ jqSecurityContext.Count()
}
