// OPR-R10-RBAC - Runs as Cluster Admin
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

	if strings.Contains(fmt.Sprintf("%v", jqCRB), "cluster-admin") {
		rbac++
	}

	return rbac

}
