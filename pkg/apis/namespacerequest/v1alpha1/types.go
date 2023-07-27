package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceRequestState string

const (
	NamespaceStatusPending NamespaceRequestState = "Pending"
	NamespaceStatusCreated NamespaceRequestState = "Created"
	NamespaceStatusError   NamespaceRequestState = "Error"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespaceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceRequestSpec   `json:"spec"`
	Status NamespaceRequestStatus `json:"status"`
}

type NamespaceRequestSpec struct {
	NamespaceName string `json:"namespacename"`
	Approved      bool   `json:"approved"`
}

type NamespaceRequestStatus struct {
	State   NamespaceRequestState `json:"state"`
	Message string                `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespaceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NamespaceRequest `json:"items"`
}
