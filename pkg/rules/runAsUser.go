// OPR-R8-SC - securityContext set to runAsUser: 0
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func RunAsUser(json []byte) int {
	sc := 0
	spec := getSpecSelector(json)

	jqContainers := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".containers").
		Where("securityContext", "!=", nil).
		Where("securityContext.runAsUser", "!=", nil).
		Where("securityContext.runAsUser", "=", 0)

	jqSecurityContext := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".securityContext").
		Where("securityContext", "!=", nil).
		Where("securityContext.privileged", "!=", nil)

	if strings.Contains(fmt.Sprintf("%v", jqSecurityContext.Get()), "runAsUser:0") {
		sc++
	}

	return jqContainers.Count() + sc
}
