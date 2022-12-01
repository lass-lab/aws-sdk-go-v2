// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/pipes/types"
)

func ExampleMQBrokerAccessCredentials_outputUsage() {
	var union types.MQBrokerAccessCredentials
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MQBrokerAccessCredentialsMemberBasicAuth:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleMSKAccessCredentials_outputUsage() {
	var union types.MSKAccessCredentials
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MSKAccessCredentialsMemberClientCertificateTlsAuth:
		_ = v.Value // Value is string

	case *types.MSKAccessCredentialsMemberSaslScram512Auth:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleSelfManagedKafkaAccessConfigurationCredentials_outputUsage() {
	var union types.SelfManagedKafkaAccessConfigurationCredentials
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SelfManagedKafkaAccessConfigurationCredentialsMemberBasicAuth:
		_ = v.Value // Value is string

	case *types.SelfManagedKafkaAccessConfigurationCredentialsMemberClientCertificateTlsAuth:
		_ = v.Value // Value is string

	case *types.SelfManagedKafkaAccessConfigurationCredentialsMemberSaslScram256Auth:
		_ = v.Value // Value is string

	case *types.SelfManagedKafkaAccessConfigurationCredentialsMemberSaslScram512Auth:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
