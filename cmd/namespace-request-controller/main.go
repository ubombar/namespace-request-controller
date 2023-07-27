package main

import (
	"fmt"

	"github.com/ubombar/namespace-request-controller/pkg/apis/namespacerequest/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	fmt.Println("Test")

	a := v1alpha1.NamespaceRequest{
		ObjectMeta: v1.ObjectMeta{
			Name:      "test",
			Namespace: "Tets",
		},
		Spec: v1alpha1.NamespaceRequestSpec{
			NamespaceName: "test-namespace",
		},
	}

	fmt.Printf("a: %v\n", a)
}
