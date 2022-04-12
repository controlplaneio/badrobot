// OPR-R9-RBAC - Runs as Cluster Admin
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ClusterAdmin(json []byte) int {
	rbac := 0

	jqCRB := gojsonq.New().Reader(bytes.NewReader(json)).
		From("roleRef.name").Get()

	fmt.Printf("%v", jqCRB)

	if strings.Contains(fmt.Sprintf("%v", jqCRB), "cluster-admin") {
		rbac++
	}

	fmt.Printf("%v", rbac)
	return rbac

}
