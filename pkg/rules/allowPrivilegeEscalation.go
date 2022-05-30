// OPR-R4-SC - securityContext set to allowPrivilegeEscalation: true
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func AllowPrivilegeEscalation(json []byte) (int, error) {
	sc := 0
	spec := getSpecSelector(json)

	jqContainers := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".containers").
		Where("securityContext", "!=", nil).
		Where("securityContext.allowPrivilegeEscalation", "!=", nil).
		Where("securityContext.allowPrivilegeEscalation", "=", true)

	jqSecurityContext := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".securityContext").
		Where("securityContext", "!=", nil).
		Where("securityContext.allowPrivilegeEscalation", "!=", nil)

	if strings.Contains(fmt.Sprintf("%v", jqSecurityContext.Get()), "allowPrivilegeEscalation:true") {
		sc++
	}

	return jqContainers.Count() + sc, nil
}
