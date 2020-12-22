// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CloudWatchEventSource string

// Values returns all known values for CloudWatchEventSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CloudWatchEventSource) Values() []CloudWatchEventSource {
	return []CloudWatchEventSource{
		"EC2",
		"CODE_DEPLOY",
		"HEALTH",
		"RDS",
	}
}

type ConfigurationEventResourceType string

// Values returns all known values for ConfigurationEventResourceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConfigurationEventResourceType) Values() []ConfigurationEventResourceType {
	return []ConfigurationEventResourceType{
		"CLOUDWATCH_ALARM",
		"CLOUDWATCH_LOG",
		"CLOUDFORMATION",
		"SSM_ASSOCIATION",
	}
}

type ConfigurationEventStatus string

// Values returns all known values for ConfigurationEventStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationEventStatus) Values() []ConfigurationEventStatus {
	return []ConfigurationEventStatus{
		"INFO",
		"WARN",
		"ERROR",
	}
}

type FeedbackKey string

// Values returns all known values for FeedbackKey. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FeedbackKey) Values() []FeedbackKey {
	return []FeedbackKey{
		"INSIGHTS_FEEDBACK",
	}
}

type FeedbackValue string

// Values returns all known values for FeedbackValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FeedbackValue) Values() []FeedbackValue {
	return []FeedbackValue{
		"NOT_SPECIFIED",
		"USEFUL",
		"NOT_USEFUL",
	}
}

type LogFilter string

// Values returns all known values for LogFilter. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogFilter) Values() []LogFilter {
	return []LogFilter{
		"ERROR",
		"WARN",
		"INFO",
	}
}

type OsType string

// Values returns all known values for OsType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OsType) Values() []OsType {
	return []OsType{
		"WINDOWS",
		"LINUX",
	}
}

type SeverityLevel string

// Values returns all known values for SeverityLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SeverityLevel) Values() []SeverityLevel {
	return []SeverityLevel{
		"Low",
		"Medium",
		"High",
	}
}

type Status string

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"IGNORE",
		"RESOLVED",
		"PENDING",
	}
}

type Tier string

// Values returns all known values for Tier. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Tier) Values() []Tier {
	return []Tier{
		"CUSTOM",
		"DEFAULT",
		"DOT_NET_CORE",
		"DOT_NET_WORKER",
		"DOT_NET_WEB_TIER",
		"DOT_NET_WEB",
		"SQL_SERVER",
		"SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP",
		"MYSQL",
		"POSTGRESQL",
		"JAVA_JMX",
		"ORACLE",
	}
}
