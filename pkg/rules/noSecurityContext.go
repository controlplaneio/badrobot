package rules

import (
	"bytes"

	"github.com/thedevsaddam/gojsonq/v2"
)

func NoSecurityContext(json []byte) int {
	spec := getSpecSelector(json)
	sc := 0

	jqContainers := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec + ".containers").
		Select("securityContext")

	jqSecurityContext := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec + ".securityContext")

	if jqContainers.Count() == 0 && jqSecurityContext.Count() == 0 {
		sc++
	}

	return sc
}
