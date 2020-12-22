// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the traffic policy instances that you created by using
// the current AWS account. After you submit an UpdateTrafficPolicyInstance
// request, there's a brief delay while Amazon Route 53 creates the resource record
// sets that are specified in the traffic policy definition. For more information,
// see the State response element. Route 53 returns a maximum of 100 items in each
// response. If you have a lot of traffic policy instances, you can use the
// MaxItems parameter to list them in groups of up to 100.
func (c *Client) ListTrafficPolicyInstances(ctx context.Context, params *ListTrafficPolicyInstancesInput, optFns ...func(*Options)) (*ListTrafficPolicyInstancesOutput, error) {
	if params == nil {
		params = &ListTrafficPolicyInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrafficPolicyInstances", params, optFns, addOperationListTrafficPolicyInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrafficPolicyInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about the traffic policy instances that you created
// by using the current AWS account.
type ListTrafficPolicyInstancesInput struct {

	// If the value of IsTruncated in the previous response was true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstances request. For the value of HostedZoneId, specify the
	// value of HostedZoneIdMarker from the previous response, which is the hosted zone
	// ID of the first traffic policy instance in the next group of traffic policy
	// instances. If the value of IsTruncated in the previous response was false, there
	// are no more traffic policy instances to get.
	HostedZoneIdMarker *string

	// The maximum number of traffic policy instances that you want Amazon Route 53 to
	// return in response to a ListTrafficPolicyInstances request. If you have more
	// than MaxItems traffic policy instances, the value of the IsTruncated element in
	// the response is true, and the values of HostedZoneIdMarker,
	// TrafficPolicyInstanceNameMarker, and TrafficPolicyInstanceTypeMarker represent
	// the first traffic policy instance in the next group of MaxItems traffic policy
	// instances.
	MaxItems *int32

	// If the value of IsTruncated in the previous response was true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstances request. For the value of trafficpolicyinstancename,
	// specify the value of TrafficPolicyInstanceNameMarker from the previous response,
	// which is the name of the first traffic policy instance in the next group of
	// traffic policy instances. If the value of IsTruncated in the previous response
	// was false, there are no more traffic policy instances to get.
	TrafficPolicyInstanceNameMarker *string

	// If the value of IsTruncated in the previous response was true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstances request. For the value of trafficpolicyinstancetype,
	// specify the value of TrafficPolicyInstanceTypeMarker from the previous response,
	// which is the type of the first traffic policy instance in the next group of
	// traffic policy instances. If the value of IsTruncated in the previous response
	// was false, there are no more traffic policy instances to get.
	TrafficPolicyInstanceTypeMarker types.RRType
}

// A complex type that contains the response information for the request.
type ListTrafficPolicyInstancesOutput struct {

	// A flag that indicates whether there are more traffic policy instances to be
	// listed. If the response was truncated, you can get more traffic policy instances
	// by calling ListTrafficPolicyInstances again and specifying the values of the
	// HostedZoneIdMarker, TrafficPolicyInstanceNameMarker, and
	// TrafficPolicyInstanceTypeMarker in the corresponding request parameters.
	//
	// This member is required.
	IsTruncated bool

	// The value that you specified for the MaxItems parameter in the call to
	// ListTrafficPolicyInstances that produced the current response.
	//
	// This member is required.
	MaxItems *int32

	// A list that contains one TrafficPolicyInstance element for each traffic policy
	// instance that matches the elements in the request.
	//
	// This member is required.
	TrafficPolicyInstances []types.TrafficPolicyInstance

	// If IsTruncated is true, HostedZoneIdMarker is the ID of the hosted zone of the
	// first traffic policy instance that Route 53 will return if you submit another
	// ListTrafficPolicyInstances request.
	HostedZoneIdMarker *string

	// If IsTruncated is true, TrafficPolicyInstanceNameMarker is the name of the first
	// traffic policy instance that Route 53 will return if you submit another
	// ListTrafficPolicyInstances request.
	TrafficPolicyInstanceNameMarker *string

	// If IsTruncated is true, TrafficPolicyInstanceTypeMarker is the DNS type of the
	// resource record sets that are associated with the first traffic policy instance
	// that Amazon Route 53 will return if you submit another
	// ListTrafficPolicyInstances request.
	TrafficPolicyInstanceTypeMarker types.RRType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTrafficPolicyInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTrafficPolicyInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTrafficPolicyInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficPolicyInstances(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListTrafficPolicyInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListTrafficPolicyInstances",
	}
}
