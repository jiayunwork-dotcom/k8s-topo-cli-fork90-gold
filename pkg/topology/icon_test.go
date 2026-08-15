package topology

import "testing"

func TestStatusIcon_Degraded(t *testing.T) {
	got := StatusIcon("Degraded")
	if got != "🟡" {
		t.Fatalf("StatusIcon(Degraded)=%q, want 🟡", got)
	}
}
