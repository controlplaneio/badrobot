// OPR-R6-SC - securityContext set to readOnlyRootFilesystem: false
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ReadOnlyRootFilesystem(json []byte) int {
	sc := 0
	spec := getSpecSelector(json)

	jqContainers := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".containers").
		Where("securityContext", "!=", nil).
		Where("securityContext.readOnlyRootFilesystem", "!=", nil).
		Where("securityContext.readOnlyRootFilesystem", "=", false)

	jqSecurityContext := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".securityContext").
		Where("securityContext", "!=", nil).
		Where("securityContext.readOnlyRootFilesystem", "!=", nil)

	if strings.Contains(fmt.Sprintf("%v", jqSecurityContext.Get()), "readOnlyRootFilesystem:false") {
		sc++
	}

	return jqContainers.Count() + sc
}
