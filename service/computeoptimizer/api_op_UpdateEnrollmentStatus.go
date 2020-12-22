// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the enrollment (opt in) status of an account to the AWS Compute
// Optimizer service. If the account is a management account of an organization,
// this action can also be used to enroll member accounts within the organization.
func (c *Client) UpdateEnrollmentStatus(ctx context.Context, params *UpdateEnrollmentStatusInput, optFns ...func(*Options)) (*UpdateEnrollmentStatusOutput, error) {
	if params == nil {
		params = &UpdateEnrollmentStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnrollmentStatus", params, optFns, addOperationUpdateEnrollmentStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnrollmentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnrollmentStatusInput struct {

	// The new enrollment status of the account. Accepted options are Active or
	// Inactive. You will get an error if Pending or Failed are specified.
	//
	// This member is required.
	Status types.Status

	// Indicates whether to enroll member accounts of the organization if the your
	// account is the management account of an organization.
	IncludeMemberAccounts bool
}

type UpdateEnrollmentStatusOutput struct {

	// The enrollment status of the account.
	Status types.Status

	// The reason for the enrollment status of the account. For example, an account
	// might show a status of Pending because member accounts of an organization
	// require more time to be enrolled in the service.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateEnrollmentStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateEnrollmentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateEnrollmentStatus{}, middleware.After)
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
	if err = addOpUpdateEnrollmentStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnrollmentStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateEnrollmentStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "UpdateEnrollmentStatus",
	}
}
