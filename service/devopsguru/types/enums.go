// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AnomalySeverity string

// Enum values for AnomalySeverity
const (
	AnomalySeverityLow    AnomalySeverity = "LOW"
	AnomalySeverityMedium AnomalySeverity = "MEDIUM"
	AnomalySeverityHigh   AnomalySeverity = "HIGH"
)

// Values returns all known values for AnomalySeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnomalySeverity) Values() []AnomalySeverity {
	return []AnomalySeverity{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

type AnomalyStatus string

// Enum values for AnomalyStatus
const (
	AnomalyStatusOngoing AnomalyStatus = "ONGOING"
	AnomalyStatusClosed  AnomalyStatus = "CLOSED"
)

// Values returns all known values for AnomalyStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnomalyStatus) Values() []AnomalyStatus {
	return []AnomalyStatus{
		"ONGOING",
		"CLOSED",
	}
}

type CloudWatchMetricsStat string

// Enum values for CloudWatchMetricsStat
const (
	CloudWatchMetricsStatSum         CloudWatchMetricsStat = "Sum"
	CloudWatchMetricsStatAverage     CloudWatchMetricsStat = "Average"
	CloudWatchMetricsStatSamplecount CloudWatchMetricsStat = "SampleCount"
	CloudWatchMetricsStatMinimum     CloudWatchMetricsStat = "Minimum"
	CloudWatchMetricsStatMaximum     CloudWatchMetricsStat = "Maximum"
	CloudWatchMetricsStatP99         CloudWatchMetricsStat = "p99"
	CloudWatchMetricsStatP90         CloudWatchMetricsStat = "p90"
	CloudWatchMetricsStatP50         CloudWatchMetricsStat = "p50"
)

// Values returns all known values for CloudWatchMetricsStat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CloudWatchMetricsStat) Values() []CloudWatchMetricsStat {
	return []CloudWatchMetricsStat{
		"Sum",
		"Average",
		"SampleCount",
		"Minimum",
		"Maximum",
		"p99",
		"p90",
		"p50",
	}
}

type EventClass string

// Enum values for EventClass
const (
	EventClassInfrastructure EventClass = "INFRASTRUCTURE"
	EventClassDeployment     EventClass = "DEPLOYMENT"
	EventClassSecurityChange EventClass = "SECURITY_CHANGE"
	EventClassConfigChange   EventClass = "CONFIG_CHANGE"
	EventClassSchemaChange   EventClass = "SCHEMA_CHANGE"
)

// Values returns all known values for EventClass. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EventClass) Values() []EventClass {
	return []EventClass{
		"INFRASTRUCTURE",
		"DEPLOYMENT",
		"SECURITY_CHANGE",
		"CONFIG_CHANGE",
		"SCHEMA_CHANGE",
	}
}

type EventDataSource string

// Enum values for EventDataSource
const (
	EventDataSourceAwsCloudTrail EventDataSource = "AWS_CLOUD_TRAIL"
	EventDataSourceAwsCodeDeploy EventDataSource = "AWS_CODE_DEPLOY"
)

// Values returns all known values for EventDataSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EventDataSource) Values() []EventDataSource {
	return []EventDataSource{
		"AWS_CLOUD_TRAIL",
		"AWS_CODE_DEPLOY",
	}
}

type InsightFeedbackOption string

// Enum values for InsightFeedbackOption
const (
	InsightFeedbackOptionValidCollection      InsightFeedbackOption = "VALID_COLLECTION"
	InsightFeedbackOptionRecommendationUseful InsightFeedbackOption = "RECOMMENDATION_USEFUL"
	InsightFeedbackOptionAlertTooSensitive    InsightFeedbackOption = "ALERT_TOO_SENSITIVE"
	InsightFeedbackOptionDataNoisyAnomaly     InsightFeedbackOption = "DATA_NOISY_ANOMALY"
	InsightFeedbackOptionDataIncorrect        InsightFeedbackOption = "DATA_INCORRECT"
)

// Values returns all known values for InsightFeedbackOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InsightFeedbackOption) Values() []InsightFeedbackOption {
	return []InsightFeedbackOption{
		"VALID_COLLECTION",
		"RECOMMENDATION_USEFUL",
		"ALERT_TOO_SENSITIVE",
		"DATA_NOISY_ANOMALY",
		"DATA_INCORRECT",
	}
}

type InsightSeverity string

// Enum values for InsightSeverity
const (
	InsightSeverityLow    InsightSeverity = "LOW"
	InsightSeverityMedium InsightSeverity = "MEDIUM"
	InsightSeverityHigh   InsightSeverity = "HIGH"
)

// Values returns all known values for InsightSeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InsightSeverity) Values() []InsightSeverity {
	return []InsightSeverity{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

type InsightStatus string

// Enum values for InsightStatus
const (
	InsightStatusOngoing InsightStatus = "ONGOING"
	InsightStatusClosed  InsightStatus = "CLOSED"
)

// Values returns all known values for InsightStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InsightStatus) Values() []InsightStatus {
	return []InsightStatus{
		"ONGOING",
		"CLOSED",
	}
}

type InsightType string

// Enum values for InsightType
const (
	InsightTypeReactive  InsightType = "REACTIVE"
	InsightTypeProactive InsightType = "PROACTIVE"
)

// Values returns all known values for InsightType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (InsightType) Values() []InsightType {
	return []InsightType{
		"REACTIVE",
		"PROACTIVE",
	}
}

type OptInStatus string

// Enum values for OptInStatus
const (
	OptInStatusEnabled  OptInStatus = "ENABLED"
	OptInStatusDisabled OptInStatus = "DISABLED"
)

// Values returns all known values for OptInStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OptInStatus) Values() []OptInStatus {
	return []OptInStatus{
		"ENABLED",
		"DISABLED",
	}
}

type UpdateResourceCollectionAction string

// Enum values for UpdateResourceCollectionAction
const (
	UpdateResourceCollectionActionAdd    UpdateResourceCollectionAction = "ADD"
	UpdateResourceCollectionActionRemove UpdateResourceCollectionAction = "REMOVE"
)

// Values returns all known values for UpdateResourceCollectionAction. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (UpdateResourceCollectionAction) Values() []UpdateResourceCollectionAction {
	return []UpdateResourceCollectionAction{
		"ADD",
		"REMOVE",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"CANNOT_PARSE",
		"FIELD_VALIDATION_FAILED",
		"OTHER",
	}
}
