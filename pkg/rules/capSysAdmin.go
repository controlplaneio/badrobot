// OPR-R9-SC - securityContext adds CAP_SYS_ADMIN Linux capability
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func CapSysAdmin(json []byte) int {
	sc := 0
	spec := getSpecSelector(json)

	capAdd := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec + ".containers").
		Only("securityContext.capabilities.add")

	fmt.Printf("%v", capAdd)

	if capAdd != nil && strings.Contains(fmt.Sprintf("%v", capAdd), "SYS_ADMIN") {
		sc++
	}

	capAddSpec := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec + ".securityContext.capabilities.add").Get()

	fmt.Printf("%v", capAddSpec)

	if strings.Contains(fmt.Sprintf("%v", capAddSpec), "SYS_ADMIN") {
		sc++
	}

	return sc
}
