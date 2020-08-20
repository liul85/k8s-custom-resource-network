package v1

import (
	"github.com/liul85/k8s-custom-resource-network/pkg/apis/networkcrd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemaGroupVersion is the identifier for the API
var SchemaGroupVersion = schema.GroupVersion{
	Group:   networkcrd.GroupName,
	Version: networkcrd.Version,
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemaGroupVersion.WithResource(resource).GroupResource()
}

//Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemaGroupVersion.WithKind(kind).GroupKind()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemaGroupVersion, &Network{}, &NetworkList{})
	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}
