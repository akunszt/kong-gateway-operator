package konnect

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	configurationv1 "github.com/kong/kubernetes-configuration/api/configuration/v1"
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	configurationv1beta1 "github.com/kong/kubernetes-configuration/api/configuration/v1beta1"
	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
)

// SupportedKonnectEntityType is an interface that all Konnect entity types
// must implement.
type SupportedKonnectEntityType interface {
	konnectv1alpha1.KonnectControlPlane |
		configurationv1alpha1.KongService |
		configurationv1alpha1.KongRoute |
		configurationv1.KongConsumer |
		configurationv1beta1.KongConsumerGroup
	// TODO: add other types

	GetTypeName() string
}

// EntityType is an interface that all Konnect entity types must implement.
// Separating this from SupportedKonnectEntityType allows us to use EntityType
// where client.Object is required, since it embeds client.Object and uses pointer
// to refer to the SupportedKonnectEntityType.
type EntityType[T SupportedKonnectEntityType] interface {
	*T
	// Kubernetes Object methods
	GetObjectMeta() metav1.Object
	client.Object

	// Additional methods which are used in reconciling Konnect entities.
	GetConditions() []metav1.Condition
	SetConditions([]metav1.Condition)
	GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus
	SetKonnectID(string)
}

// EntityWithKonnectAPIAuthConfigurationRef is an interface that all Konnect entity types
// that reference a KonnectAPIAuthConfiguration must implement.
// More specifically Konnect's ControlPlane does implement that, while all the other
// Konnect entities that are defined within a ControlPlane do not because their
// KonnectAPIAuthConfigurationRef is defined in the referenced ControlPlane.
type EntityWithKonnectAPIAuthConfigurationRef interface {
	GetKonnectAPIAuthConfigurationRef() konnectv1alpha1.KonnectAPIAuthConfigurationRef
}
