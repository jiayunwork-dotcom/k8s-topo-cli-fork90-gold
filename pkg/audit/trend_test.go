package audit

import "testing"

func TestDiffAuditReports_SmallIncreaseTrendUp(t *testing.T) {
	v := func(val string) NamespaceAuditResult {
		return NamespaceAuditResult{
			Namespace: "ns",
			Violations: []Violation{{
				Namespace:    "ns",
				Dimension:    DimMaxPods,
				CurrentValue: val,
			}},
		}
	}
	diff := DiffAuditReports(
		&AuditReport{Namespaces: []NamespaceAuditResult{v("106")}},
		&AuditReport{Namespaces: []NamespaceAuditResult{v("100")}},
	)
	if len(diff.Trends) != 1 {
		t.Fatalf("got %d trends, want 1", len(diff.Trends))
	}
	if diff.Trends[0].Direction != TrendUp {
		t.Fatalf("6%% increase Direction=%q, want %q", diff.Trends[0].Direction, TrendUp)
	}
}
