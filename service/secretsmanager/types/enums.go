// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type FilterNameStringType string

// Values returns all known values for FilterNameStringType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterNameStringType) Values() []FilterNameStringType {
	return []FilterNameStringType{
		"description",
		"name",
		"tag-key",
		"tag-value",
		"all",
	}
}

type SortOrderType string

// Enum values for SortOrderType
const (
	SortOrderTypeAsc  SortOrderType = "asc"
	SortOrderTypeDesc SortOrderType = "desc"
)

// Values returns all known values for SortOrderType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SortOrderType) Values() []SortOrderType {
	return []SortOrderType{
		"asc",
		"desc",
	}
}
