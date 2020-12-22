// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthMechanismValue string

// Enum values for AuthMechanismValue
const (
	AuthMechanismValueDefault   AuthMechanismValue = "default"
	AuthMechanismValueMongodbCr AuthMechanismValue = "mongodb_cr"
	AuthMechanismValueScramSha1 AuthMechanismValue = "scram_sha_1"
)

// Values returns all known values for AuthMechanismValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthMechanismValue) Values() []AuthMechanismValue {
	return []AuthMechanismValue{
		"default",
		"mongodb_cr",
		"scram_sha_1",
	}
}

type AuthTypeValue string

// Enum values for AuthTypeValue
const (
	AuthTypeValueNo       AuthTypeValue = "no"
	AuthTypeValuePassword AuthTypeValue = "password"
)

// Values returns all known values for AuthTypeValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthTypeValue) Values() []AuthTypeValue {
	return []AuthTypeValue{
		"no",
		"password",
	}
}

type CharLengthSemantics string

// Enum values for CharLengthSemantics
const (
	CharLengthSemanticsDefault CharLengthSemantics = "default"
	CharLengthSemanticsChar    CharLengthSemantics = "char"
	CharLengthSemanticsByte    CharLengthSemantics = "byte"
)

// Values returns all known values for CharLengthSemantics. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CharLengthSemantics) Values() []CharLengthSemantics {
	return []CharLengthSemantics{
		"default",
		"char",
		"byte",
	}
}

type CompressionTypeValue string

// Enum values for CompressionTypeValue
const (
	CompressionTypeValueNone CompressionTypeValue = "none"
	CompressionTypeValueGzip CompressionTypeValue = "gzip"
)

// Values returns all known values for CompressionTypeValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CompressionTypeValue) Values() []CompressionTypeValue {
	return []CompressionTypeValue{
		"none",
		"gzip",
	}
}

type DataFormatValue string

// Enum values for DataFormatValue
const (
	DataFormatValueCsv     DataFormatValue = "csv"
	DataFormatValueParquet DataFormatValue = "parquet"
)

// Values returns all known values for DataFormatValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataFormatValue) Values() []DataFormatValue {
	return []DataFormatValue{
		"csv",
		"parquet",
	}
}

type DatePartitionDelimiterValue string

// Enum values for DatePartitionDelimiterValue
const (
	DatePartitionDelimiterValueSlash      DatePartitionDelimiterValue = "SLASH"
	DatePartitionDelimiterValueUnderscore DatePartitionDelimiterValue = "UNDERSCORE"
	DatePartitionDelimiterValueDash       DatePartitionDelimiterValue = "DASH"
	DatePartitionDelimiterValueNone       DatePartitionDelimiterValue = "NONE"
)

// Values returns all known values for DatePartitionDelimiterValue. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DatePartitionDelimiterValue) Values() []DatePartitionDelimiterValue {
	return []DatePartitionDelimiterValue{
		"SLASH",
		"UNDERSCORE",
		"DASH",
		"NONE",
	}
}

type DatePartitionSequenceValue string

// Enum values for DatePartitionSequenceValue
const (
	DatePartitionSequenceValueYyyymmdd   DatePartitionSequenceValue = "YYYYMMDD"
	DatePartitionSequenceValueYyyymmddhh DatePartitionSequenceValue = "YYYYMMDDHH"
	DatePartitionSequenceValueYyyymm     DatePartitionSequenceValue = "YYYYMM"
	DatePartitionSequenceValueMmyyyydd   DatePartitionSequenceValue = "MMYYYYDD"
	DatePartitionSequenceValueDdmmyyyy   DatePartitionSequenceValue = "DDMMYYYY"
)

// Values returns all known values for DatePartitionSequenceValue. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DatePartitionSequenceValue) Values() []DatePartitionSequenceValue {
	return []DatePartitionSequenceValue{
		"YYYYMMDD",
		"YYYYMMDDHH",
		"YYYYMM",
		"MMYYYYDD",
		"DDMMYYYY",
	}
}

type DmsSslModeValue string

// Enum values for DmsSslModeValue
const (
	DmsSslModeValueNone       DmsSslModeValue = "none"
	DmsSslModeValueRequire    DmsSslModeValue = "require"
	DmsSslModeValueVerifyCa   DmsSslModeValue = "verify-ca"
	DmsSslModeValueVerifyFull DmsSslModeValue = "verify-full"
)

// Values returns all known values for DmsSslModeValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DmsSslModeValue) Values() []DmsSslModeValue {
	return []DmsSslModeValue{
		"none",
		"require",
		"verify-ca",
		"verify-full",
	}
}

type EncodingTypeValue string

// Enum values for EncodingTypeValue
const (
	EncodingTypeValuePlain           EncodingTypeValue = "plain"
	EncodingTypeValuePlainDictionary EncodingTypeValue = "plain-dictionary"
	EncodingTypeValueRleDictionary   EncodingTypeValue = "rle-dictionary"
)

// Values returns all known values for EncodingTypeValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncodingTypeValue) Values() []EncodingTypeValue {
	return []EncodingTypeValue{
		"plain",
		"plain-dictionary",
		"rle-dictionary",
	}
}

type EncryptionModeValue string

// Enum values for EncryptionModeValue
const (
	EncryptionModeValueSseS3  EncryptionModeValue = "sse-s3"
	EncryptionModeValueSseKms EncryptionModeValue = "sse-kms"
)

// Values returns all known values for EncryptionModeValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionModeValue) Values() []EncryptionModeValue {
	return []EncryptionModeValue{
		"sse-s3",
		"sse-kms",
	}
}

type MessageFormatValue string

// Enum values for MessageFormatValue
const (
	MessageFormatValueJson            MessageFormatValue = "json"
	MessageFormatValueJsonUnformatted MessageFormatValue = "json-unformatted"
)

// Values returns all known values for MessageFormatValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MessageFormatValue) Values() []MessageFormatValue {
	return []MessageFormatValue{
		"json",
		"json-unformatted",
	}
}

type MigrationTypeValue string

// Enum values for MigrationTypeValue
const (
	MigrationTypeValueFullLoad       MigrationTypeValue = "full-load"
	MigrationTypeValueCdc            MigrationTypeValue = "cdc"
	MigrationTypeValueFullLoadAndCdc MigrationTypeValue = "full-load-and-cdc"
)

// Values returns all known values for MigrationTypeValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MigrationTypeValue) Values() []MigrationTypeValue {
	return []MigrationTypeValue{
		"full-load",
		"cdc",
		"full-load-and-cdc",
	}
}

type NestingLevelValue string

// Enum values for NestingLevelValue
const (
	NestingLevelValueNone NestingLevelValue = "none"
	NestingLevelValueOne  NestingLevelValue = "one"
)

// Values returns all known values for NestingLevelValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NestingLevelValue) Values() []NestingLevelValue {
	return []NestingLevelValue{
		"none",
		"one",
	}
}

type ParquetVersionValue string

// Enum values for ParquetVersionValue
const (
	ParquetVersionValueParquet10 ParquetVersionValue = "parquet-1-0"
	ParquetVersionValueParquet20 ParquetVersionValue = "parquet-2-0"
)

// Values returns all known values for ParquetVersionValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ParquetVersionValue) Values() []ParquetVersionValue {
	return []ParquetVersionValue{
		"parquet-1-0",
		"parquet-2-0",
	}
}

type RefreshSchemasStatusTypeValue string

// Enum values for RefreshSchemasStatusTypeValue
const (
	RefreshSchemasStatusTypeValueSuccessful RefreshSchemasStatusTypeValue = "successful"
	RefreshSchemasStatusTypeValueFailed     RefreshSchemasStatusTypeValue = "failed"
	RefreshSchemasStatusTypeValueRefreshing RefreshSchemasStatusTypeValue = "refreshing"
)

// Values returns all known values for RefreshSchemasStatusTypeValue. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RefreshSchemasStatusTypeValue) Values() []RefreshSchemasStatusTypeValue {
	return []RefreshSchemasStatusTypeValue{
		"successful",
		"failed",
		"refreshing",
	}
}

type ReleaseStatusValues string

// Enum values for ReleaseStatusValues
const (
	ReleaseStatusValuesBeta ReleaseStatusValues = "beta"
)

// Values returns all known values for ReleaseStatusValues. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReleaseStatusValues) Values() []ReleaseStatusValues {
	return []ReleaseStatusValues{
		"beta",
	}
}

type ReloadOptionValue string

// Enum values for ReloadOptionValue
const (
	ReloadOptionValueDataReload   ReloadOptionValue = "data-reload"
	ReloadOptionValueValidateOnly ReloadOptionValue = "validate-only"
)

// Values returns all known values for ReloadOptionValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReloadOptionValue) Values() []ReloadOptionValue {
	return []ReloadOptionValue{
		"data-reload",
		"validate-only",
	}
}

type ReplicationEndpointTypeValue string

// Enum values for ReplicationEndpointTypeValue
const (
	ReplicationEndpointTypeValueSource ReplicationEndpointTypeValue = "source"
	ReplicationEndpointTypeValueTarget ReplicationEndpointTypeValue = "target"
)

// Values returns all known values for ReplicationEndpointTypeValue. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationEndpointTypeValue) Values() []ReplicationEndpointTypeValue {
	return []ReplicationEndpointTypeValue{
		"source",
		"target",
	}
}

type SafeguardPolicy string

// Enum values for SafeguardPolicy
const (
	SafeguardPolicyRelyOnSqlServerReplicationAgent SafeguardPolicy = "rely-on-sql-server-replication-agent"
	SafeguardPolicyExclusiveAutomaticTruncation    SafeguardPolicy = "exclusive-automatic-truncation"
	SafeguardPolicySharedAutomaticTruncation       SafeguardPolicy = "shared-automatic-truncation"
)

// Values returns all known values for SafeguardPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SafeguardPolicy) Values() []SafeguardPolicy {
	return []SafeguardPolicy{
		"rely-on-sql-server-replication-agent",
		"exclusive-automatic-truncation",
		"shared-automatic-truncation",
	}
}

type SourceType string

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"replication-instance",
	}
}

type StartReplicationTaskTypeValue string

// Enum values for StartReplicationTaskTypeValue
const (
	StartReplicationTaskTypeValueStartReplication StartReplicationTaskTypeValue = "start-replication"
	StartReplicationTaskTypeValueResumeProcessing StartReplicationTaskTypeValue = "resume-processing"
	StartReplicationTaskTypeValueReloadTarget     StartReplicationTaskTypeValue = "reload-target"
)

// Values returns all known values for StartReplicationTaskTypeValue. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (StartReplicationTaskTypeValue) Values() []StartReplicationTaskTypeValue {
	return []StartReplicationTaskTypeValue{
		"start-replication",
		"resume-processing",
		"reload-target",
	}
}

type TargetDbType string

// Enum values for TargetDbType
const (
	TargetDbTypeSpecificDatabase  TargetDbType = "specific-database"
	TargetDbTypeMultipleDatabases TargetDbType = "multiple-databases"
)

// Values returns all known values for TargetDbType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetDbType) Values() []TargetDbType {
	return []TargetDbType{
		"specific-database",
		"multiple-databases",
	}
}
