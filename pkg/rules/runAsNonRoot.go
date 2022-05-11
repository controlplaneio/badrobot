// OPR-R7-SC - securityContext set to runAsNonRoot: false
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func RunAsNonRoot(json []byte) (int, error) {
	sc := 0
	spec := getSpecSelector(json)

	jqContainers := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".containers").
		Where("securityContext", "!=", nil).
		Where("securityContext.runAsNonRoot", "!=", nil).
		Where("securityContext.runAsNonRoot", "=", false)

	jqSecurityContext := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".securityContext").
		Where("securityContext", "!=", nil).
		Where("securityContext.privileged", "!=", nil)

	if strings.Contains(fmt.Sprintf("%v", jqSecurityContext.Get()), "runAsNonRoot:false") {
		sc++
	}

	return jqContainers.Count() + sc, nil
}
