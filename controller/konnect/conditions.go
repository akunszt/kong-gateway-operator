package konnect

// TODO(pmalek): move this to Konnect API directory so that it's part of the API contract.
// https://github.com/Kong/kubernetes-configuration/issues/14

const (
	// KonnectEntityProgrammedConditionType is the type of the condition that
	// indicates whether the entity has been programmed in Konnect.
	KonnectEntityProgrammedConditionType = "Programmed"

	// KonnectEntityProgrammedReasonProgrammed is the reason for the Programmed condition.
	// It is set when the entity has been programmed in Konnect.
	KonnectEntityProgrammedReasonProgrammed = "Programmed"
	// KonnectEntityProgrammedReasonKonnectAPIOpFailed is the reason for the Programmed condition.
	// It is set when the entity has failed to be programmed in Konnect.
	KonnectEntityProgrammedReasonKonnectAPIOpFailed = "KonnectAPIOpFailed"
)

const (
	// KonnectEntityAPIAuthConfigurationResolvedRefConditionType is the type of the
	// condition that indicates whether the APIAuth configuration reference is
	// valid and points to an existing APIAuth configuration.
	KonnectEntityAPIAuthConfigurationResolvedRefConditionType = "APIAuthResolvedRef"

	// KonnectEntityAPIAuthConfigurationResolvedRefReasonResolvedRef is the reason
	// used with the APIAuthResolvedRef condition type indicating that the APIAuth
	// configuration reference has been resolved.
	KonnectEntityAPIAuthConfigurationResolvedRefReasonResolvedRef = "ResolvedRef"
	// KonnectEntityAPIAuthConfigurationResolvedRefReasonRefNotFound is the reason
	// used with the APIAuthResolvedRef condition type indicating that the APIAuth
	// configuration reference could not be resolved.
	KonnectEntityAPIAuthConfigurationResolvedRefReasonRefNotFound = "RefNotFound"
	// KonnectEntityAPIAuthConfigurationResolvedRefReasonRefNotFound is the reason
	// used with the APIAuthResolvedRef condition type indicating that the APIAuth
	// configuration reference is invalid and could not be resolved.
	// Condition message can contain more information about the error.
	KonnectEntityAPIAuthConfigurationResolvedRefReasonRefInvalid = "RefInvalid"
)

const (
	// KonnectEntityAPIAuthConfigurationValidConditionType is the type of the
	// condition that indicates whether the referenced APIAuth configuration is
	// valid.
	KonnectEntityAPIAuthConfigurationValidConditionType = "APIAuthValid"

	// KonnectEntityAPIAuthConfigurationReasonValid is the reason used with the
	// APIAuthRefValid condition type indicating that the APIAuth configuration
	// referenced by the entity is valid.
	KonnectEntityAPIAuthConfigurationReasonValid = "Valid"
	// KonnectEntityAPIAuthConfigurationReasonInvalid is the reason used with the
	// APIAuthRefValid condition type indicating that the APIAuth configuration
	// referenced by the entity is invalid.
	KonnectEntityAPIAuthConfigurationReasonInvalid = "Invalid"
)

const (
	// ControlPlaneRefValidConditionType is the type of the condition that indicates
	// whether the ControlPlane reference is valid and points to an existing
	// ControlPlane.
	ControlPlaneRefValidConditionType = "ControlPlaneRefValid"

	// ControlPlaneRefReasonValid is the reason used with the ControlPlaneRefValid
	// condition type indicating that the ControlPlane reference is valid.
	ControlPlaneRefReasonValid = "Valid"
	// ControlPlaneRefReasonInvalid is the reason used with the ControlPlaneRefValid
	// condition type indicating that the ControlPlane reference is invalid.
	ControlPlaneRefReasonInvalid = "Invalid"
)

const (
	// KongServiceRefValidConditionType is the type of the condition that indicates
	// whether the KongService reference is valid and points to an existing
	// KongService.
	KongServiceRefValidConditionType = "KongServiceRefValid"

	// KongServiceRefReasonValid is the reason used with the KongServiceRefValid
	// condition type indicating that the KongService reference is valid.
	KongServiceRefReasonValid = "Valid"
	// KongServiceRefReasonInvalid is the reason used with the KongServiceRefValid
	// condition type indicating that the KongService reference is invalid.
	KongServiceRefReasonInvalid = "Invalid"
)
