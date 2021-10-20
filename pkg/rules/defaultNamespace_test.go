package rules

import (
	"github.com/ghodss/yaml"
	"testing"
)

// -- TODO KGW Write tests

// func Test_AllowedNamespaces(t *testing.T) {
// 	var data = `
// ---
// apiVersion: apps/v1
// kind: Deployment
// spec:
//   template:
//     spec:
//       containers:
//         - name: c1
//           securityContext:
//             allowPrivilegeEscalation: true
//         - name: c2
//           securityContext:
//             allowPrivilegeEscalation: true
//         - name: c3
//           securityContext:
//             allowPrivilegeEscalation: true
// `

// 	json, err := yaml.YAMLToJSON([]byte(data))
// 	if err != nil {
// 		t.Fatal(err.Error())
// 	}

// 	containers := AllowPrivilegeEscalation(json)
// 	if containers != 3 {
// 		t.Errorf("Got %v containers wanted %v", containers, 3)
// 	}
// }
