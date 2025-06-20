// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DeletionMode string

// Enum values for DeletionMode
const (
	DeletionModeAllConfigurations      DeletionMode = "ALL_CONFIGURATIONS"
	DeletionModeLocalConfigurationOnly DeletionMode = "LOCAL_CONFIGURATION_ONLY"
)

// Values returns all known values for DeletionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeletionMode) Values() []DeletionMode {
	return []DeletionMode{
		"ALL_CONFIGURATIONS",
		"LOCAL_CONFIGURATION_ONLY",
	}
}

type IpAddressType string

// Enum values for IpAddressType
const (
	IpAddressTypeIpv4Only  IpAddressType = "IPV4_ONLY"
	IpAddressTypeIpv6Only  IpAddressType = "IPV6_ONLY"
	IpAddressTypeDualStack IpAddressType = "DUAL_STACK"
)

// Values returns all known values for IpAddressType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IpAddressType) Values() []IpAddressType {
	return []IpAddressType{
		"IPV4_ONLY",
		"IPV6_ONLY",
		"DUAL_STACK",
	}
}

type LifeCycleState string

// Enum values for LifeCycleState
const (
	LifeCycleStateCreating  LifeCycleState = "creating"
	LifeCycleStateAvailable LifeCycleState = "available"
	LifeCycleStateUpdating  LifeCycleState = "updating"
	LifeCycleStateDeleting  LifeCycleState = "deleting"
	LifeCycleStateDeleted   LifeCycleState = "deleted"
	LifeCycleStateError     LifeCycleState = "error"
)

// Values returns all known values for LifeCycleState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LifeCycleState) Values() []LifeCycleState {
	return []LifeCycleState{
		"creating",
		"available",
		"updating",
		"deleting",
		"deleted",
		"error",
	}
}

type PerformanceMode string

// Enum values for PerformanceMode
const (
	PerformanceModeGeneralPurpose PerformanceMode = "generalPurpose"
	PerformanceModeMaxIo          PerformanceMode = "maxIO"
)

// Values returns all known values for PerformanceMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PerformanceMode) Values() []PerformanceMode {
	return []PerformanceMode{
		"generalPurpose",
		"maxIO",
	}
}

type ReplicationOverwriteProtection string

// Enum values for ReplicationOverwriteProtection
const (
	ReplicationOverwriteProtectionEnabled     ReplicationOverwriteProtection = "ENABLED"
	ReplicationOverwriteProtectionDisabled    ReplicationOverwriteProtection = "DISABLED"
	ReplicationOverwriteProtectionReplicating ReplicationOverwriteProtection = "REPLICATING"
)

// Values returns all known values for ReplicationOverwriteProtection. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationOverwriteProtection) Values() []ReplicationOverwriteProtection {
	return []ReplicationOverwriteProtection{
		"ENABLED",
		"DISABLED",
		"REPLICATING",
	}
}

type ReplicationStatus string

// Enum values for ReplicationStatus
const (
	ReplicationStatusEnabled  ReplicationStatus = "ENABLED"
	ReplicationStatusEnabling ReplicationStatus = "ENABLING"
	ReplicationStatusDeleting ReplicationStatus = "DELETING"
	ReplicationStatusError    ReplicationStatus = "ERROR"
	ReplicationStatusPaused   ReplicationStatus = "PAUSED"
	ReplicationStatusPausing  ReplicationStatus = "PAUSING"
)

// Values returns all known values for ReplicationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationStatus) Values() []ReplicationStatus {
	return []ReplicationStatus{
		"ENABLED",
		"ENABLING",
		"DELETING",
		"ERROR",
		"PAUSED",
		"PAUSING",
	}
}

type Resource string

// Enum values for Resource
const (
	ResourceFileSystem  Resource = "FILE_SYSTEM"
	ResourceMountTarget Resource = "MOUNT_TARGET"
)

// Values returns all known values for Resource. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Resource) Values() []Resource {
	return []Resource{
		"FILE_SYSTEM",
		"MOUNT_TARGET",
	}
}

type ResourceIdType string

// Enum values for ResourceIdType
const (
	ResourceIdTypeLongId  ResourceIdType = "LONG_ID"
	ResourceIdTypeShortId ResourceIdType = "SHORT_ID"
)

// Values returns all known values for ResourceIdType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceIdType) Values() []ResourceIdType {
	return []ResourceIdType{
		"LONG_ID",
		"SHORT_ID",
	}
}

type Status string

// Enum values for Status
const (
	StatusEnabled   Status = "ENABLED"
	StatusEnabling  Status = "ENABLING"
	StatusDisabled  Status = "DISABLED"
	StatusDisabling Status = "DISABLING"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"ENABLED",
		"ENABLING",
		"DISABLED",
		"DISABLING",
	}
}

type ThroughputMode string

// Enum values for ThroughputMode
const (
	ThroughputModeBursting    ThroughputMode = "bursting"
	ThroughputModeProvisioned ThroughputMode = "provisioned"
	ThroughputModeElastic     ThroughputMode = "elastic"
)

// Values returns all known values for ThroughputMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThroughputMode) Values() []ThroughputMode {
	return []ThroughputMode{
		"bursting",
		"provisioned",
		"elastic",
	}
}

type TransitionToArchiveRules string

// Enum values for TransitionToArchiveRules
const (
	TransitionToArchiveRulesAfter1Day    TransitionToArchiveRules = "AFTER_1_DAY"
	TransitionToArchiveRulesAfter7Days   TransitionToArchiveRules = "AFTER_7_DAYS"
	TransitionToArchiveRulesAfter14Days  TransitionToArchiveRules = "AFTER_14_DAYS"
	TransitionToArchiveRulesAfter30Days  TransitionToArchiveRules = "AFTER_30_DAYS"
	TransitionToArchiveRulesAfter60Days  TransitionToArchiveRules = "AFTER_60_DAYS"
	TransitionToArchiveRulesAfter90Days  TransitionToArchiveRules = "AFTER_90_DAYS"
	TransitionToArchiveRulesAfter180Days TransitionToArchiveRules = "AFTER_180_DAYS"
	TransitionToArchiveRulesAfter270Days TransitionToArchiveRules = "AFTER_270_DAYS"
	TransitionToArchiveRulesAfter365Days TransitionToArchiveRules = "AFTER_365_DAYS"
)

// Values returns all known values for TransitionToArchiveRules. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TransitionToArchiveRules) Values() []TransitionToArchiveRules {
	return []TransitionToArchiveRules{
		"AFTER_1_DAY",
		"AFTER_7_DAYS",
		"AFTER_14_DAYS",
		"AFTER_30_DAYS",
		"AFTER_60_DAYS",
		"AFTER_90_DAYS",
		"AFTER_180_DAYS",
		"AFTER_270_DAYS",
		"AFTER_365_DAYS",
	}
}

type TransitionToIARules string

// Enum values for TransitionToIARules
const (
	TransitionToIARulesAfter7Days   TransitionToIARules = "AFTER_7_DAYS"
	TransitionToIARulesAfter14Days  TransitionToIARules = "AFTER_14_DAYS"
	TransitionToIARulesAfter30Days  TransitionToIARules = "AFTER_30_DAYS"
	TransitionToIARulesAfter60Days  TransitionToIARules = "AFTER_60_DAYS"
	TransitionToIARulesAfter90Days  TransitionToIARules = "AFTER_90_DAYS"
	TransitionToIARulesAfter1Day    TransitionToIARules = "AFTER_1_DAY"
	TransitionToIARulesAfter180Days TransitionToIARules = "AFTER_180_DAYS"
	TransitionToIARulesAfter270Days TransitionToIARules = "AFTER_270_DAYS"
	TransitionToIARulesAfter365Days TransitionToIARules = "AFTER_365_DAYS"
)

// Values returns all known values for TransitionToIARules. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TransitionToIARules) Values() []TransitionToIARules {
	return []TransitionToIARules{
		"AFTER_7_DAYS",
		"AFTER_14_DAYS",
		"AFTER_30_DAYS",
		"AFTER_60_DAYS",
		"AFTER_90_DAYS",
		"AFTER_1_DAY",
		"AFTER_180_DAYS",
		"AFTER_270_DAYS",
		"AFTER_365_DAYS",
	}
}

type TransitionToPrimaryStorageClassRules string

// Enum values for TransitionToPrimaryStorageClassRules
const (
	TransitionToPrimaryStorageClassRulesAfter1Access TransitionToPrimaryStorageClassRules = "AFTER_1_ACCESS"
)

// Values returns all known values for TransitionToPrimaryStorageClassRules. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TransitionToPrimaryStorageClassRules) Values() []TransitionToPrimaryStorageClassRules {
	return []TransitionToPrimaryStorageClassRules{
		"AFTER_1_ACCESS",
	}
}
