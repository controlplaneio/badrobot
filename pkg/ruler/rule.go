package ruler

import (
	"bytes"
	"fmt"

	"github.com/thedevsaddam/gojsonq/v2"
)

type NotSupportedError struct {
	Kind string
}

func (e *NotSupportedError) Error() string {
	return fmt.Sprintf("rule does not apply to kind %s", e.Kind)
}

type Rule struct {
	Selector  string
	ID        string
	Title     string
	Reason    string
	Link      string
	Kinds     []string
	Points    int
	Weight    int
	Advise    int
	Predicate func([]byte) (int, error)
}

// Eval executes the predicate if the kind matches the rule
func (r *Rule) Eval(json []byte) (int, error) {
	jq := gojsonq.New().Reader(bytes.NewReader(json)).From("kind")
	if jq.Error() != nil {
		return 0, jq.Error()
	}

	kind := fmt.Sprintf("%s", jq.Get())

	var match bool
	for _, k := range r.Kinds {
		if k == kind {
			match = true
			break
		}
	}

	if match {
		return r.Predicate(json)
	} else {
		return 0, &NotSupportedError{Kind: kind}
	}
}
