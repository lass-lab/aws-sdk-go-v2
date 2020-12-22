// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AlgorithmicStemming string

// Values returns all known values for AlgorithmicStemming. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AlgorithmicStemming) Values() []AlgorithmicStemming {
	return []AlgorithmicStemming{
		"none",
		"minimal",
		"light",
		"full",
	}
}

type AnalysisSchemeLanguage string

// Values returns all known values for AnalysisSchemeLanguage. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisSchemeLanguage) Values() []AnalysisSchemeLanguage {
	return []AnalysisSchemeLanguage{
		"ar",
		"bg",
		"ca",
		"cs",
		"da",
		"de",
		"el",
		"en",
		"es",
		"eu",
		"fa",
		"fi",
		"fr",
		"ga",
		"gl",
		"he",
		"hi",
		"hu",
		"hy",
		"id",
		"it",
		"ja",
		"ko",
		"lv",
		"mul",
		"nl",
		"no",
		"pt",
		"ro",
		"ru",
		"sv",
		"th",
		"tr",
		"zh-Hans",
		"zh-Hant",
	}
}

type IndexFieldType string

// Values returns all known values for IndexFieldType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IndexFieldType) Values() []IndexFieldType {
	return []IndexFieldType{
		"int",
		"double",
		"literal",
		"text",
		"date",
		"latlon",
		"int-array",
		"double-array",
		"literal-array",
		"text-array",
		"date-array",
	}
}

type OptionState string

// Values returns all known values for OptionState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OptionState) Values() []OptionState {
	return []OptionState{
		"RequiresIndexDocuments",
		"Processing",
		"Active",
		"FailedToValidate",
	}
}

type PartitionInstanceType string

// Values returns all known values for PartitionInstanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PartitionInstanceType) Values() []PartitionInstanceType {
	return []PartitionInstanceType{
		"search.m1.small",
		"search.m1.large",
		"search.m2.xlarge",
		"search.m2.2xlarge",
		"search.m3.medium",
		"search.m3.large",
		"search.m3.xlarge",
		"search.m3.2xlarge",
	}
}

type SuggesterFuzzyMatching string

// Values returns all known values for SuggesterFuzzyMatching. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SuggesterFuzzyMatching) Values() []SuggesterFuzzyMatching {
	return []SuggesterFuzzyMatching{
		"none",
		"low",
		"high",
	}
}

type TLSSecurityPolicy string

// Enum values for TLSSecurityPolicy
const (
	TLSSecurityPolicyPolicyMinTls10201907 TLSSecurityPolicy = "Policy-Min-TLS-1-0-2019-07"
	TLSSecurityPolicyPolicyMinTls12201907 TLSSecurityPolicy = "Policy-Min-TLS-1-2-2019-07"
)

// Values returns all known values for TLSSecurityPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TLSSecurityPolicy) Values() []TLSSecurityPolicy {
	return []TLSSecurityPolicy{
		"Policy-Min-TLS-1-0-2019-07",
		"Policy-Min-TLS-1-2-2019-07",
	}
}
