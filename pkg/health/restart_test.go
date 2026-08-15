package health

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/k8s-topo-cli/pkg/discovery"
)

func TestCalculateHealth_RestartCountFive(t *testing.T) {
	res := &discovery.DiscoveredResources{
		Pods: []*corev1.Pod{{
			ObjectMeta: metav1.ObjectMeta{Name: "p", Namespace: "ns"},
			Status: corev1.PodStatus{
				Phase: corev1.PodRunning,
				ContainerStatuses: []corev1.ContainerStatus{{
					Name:         "app",
					RestartCount: 5,
				}},
			},
		}},
	}
	got := CalculateHealth(res, nil)
	for _, d := range got.Deductions {
		if d.Points == 1 && d.Resource == "Pod" {
			t.Fatalf("RestartCount=5 produced deduction %#v, want none", d)
		}
	}
}
