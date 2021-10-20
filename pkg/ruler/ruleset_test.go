package ruler

import (
	"strings"
	"testing"

	"github.com/ghodss/yaml"
	"go.uber.org/zap"
)

const schemaDir = ""

func TestRuleset_Run(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: kube-system
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	report := NewRuleset(zap.NewNop().Sugar()).generateReport("kube.yaml", json, schemaDir)

	critical := len(report.Scoring.Critical)
	if critical < 1 {
		t.Errorf("Got %v critical rules wanted many", critical)
	}

	advise := len(report.Scoring.Advise)
	if advise < 1 {
		t.Errorf("Got %v advise rules wanted many", advise)
	}

	if report.Score > 0 {
		t.Errorf("Got score %v wanted a negative value", report.Score)
	}
}
