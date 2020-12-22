// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice/types"
)

func ExampleEntitlementValue_outputUsage() {
	var union types.EntitlementValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EntitlementValueMemberBooleanValue:
		_ = v.Value // Value is bool

	case *types.EntitlementValueMemberDoubleValue:
		_ = v.Value // Value is float64

	case *types.EntitlementValueMemberIntegerValue:
		_ = v.Value // Value is int32

	case *types.EntitlementValueMemberStringValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *int32
var _ *bool
var _ *float64
