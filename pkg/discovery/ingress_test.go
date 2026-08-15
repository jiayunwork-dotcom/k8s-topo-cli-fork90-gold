package discovery

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetIngressServices_DefaultBackend(t *testing.T) {
	ing := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{Name: "web", Namespace: "ns"},
		Spec: networkingv1.IngressSpec{
			DefaultBackend: &networkingv1.IngressBackend{
				Service: &networkingv1.IngressServiceBackend{Name: "svc-web"},
			},
		},
	}
	svcs := []*corev1.Service{{
		ObjectMeta: metav1.ObjectMeta{Name: "svc-web", Namespace: "ns"},
	}}
	got := GetIngressServices(ing, svcs)
	if len(got) != 1 || got[0].Name != "svc-web" {
		t.Fatalf("GetIngressServices(defaultBackend only) got %d services, want svc-web", len(got))
	}
}
