package audit

import "testing"

func TestEvaluatePolicies_PodCountAtLimit(t *testing.T) {
	limit := int64(10)
	results := EvaluatePolicies(
		[]Policy{{
			Name:   "pods",
			Scope:  "cluster",
			Action: "warn",
			Limits: PolicyLimits{MaxPods: &limit},
		}},
		[]NamespaceStats{{Namespace: "ns", PodCount: 10}},
	)
	if len(results) != 1 {
		t.Fatalf("got %d namespace results, want 1", len(results))
	}
	if len(results[0].Violations) != 0 {
		t.Fatalf("pod count == MaxPods produced violations %#v, want none", results[0].Violations)
	}
}
