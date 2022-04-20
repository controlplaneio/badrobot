// OPR-R10-RBAC - Runs as Cluster Admin
package rules

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ClusterAdmin(json []byte) int {
	rbac := 0

	jqCRB := gojsonq.New().Reader(bytes.NewReader(json)).
		From("roleRef.name")

	reCRB := regexp.MustCompile(`:?[^-_\.](cluster):?[-](admin):?[^-_\.]`)

	if reCRB.MatchString(fmt.Sprintf("%v", jqCRB)) {
		rbac++
	}
	return rbac

}
