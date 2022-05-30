// OPR-R5-SC - securityContext set to privileged: true
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func Privileged(json []byte) (int, error) {
	sc := 0
	spec := getSpecSelector(json)

	jqContainers := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".containers").
		Where("securityContext", "!=", nil).
		Where("securityContext.privileged", "!=", nil).
		Where("securityContext.privileged", "=", true)

	jqSecurityContext := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec+".securityContext").
		Where("securityContext", "!=", nil).
		Where("securityContext.privileged", "!=", nil)

	if strings.Contains(fmt.Sprintf("%v", jqSecurityContext.Get()), "privileged:true") {
		sc++
	}

	return jqContainers.Count() + sc, nil
}
