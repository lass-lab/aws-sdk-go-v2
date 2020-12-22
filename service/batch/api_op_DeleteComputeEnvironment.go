// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an AWS Batch compute environment. Before you can delete a compute
// environment, you must set its state to DISABLED with the
// UpdateComputeEnvironment API operation and disassociate it from any job queues
// with the UpdateJobQueue API operation. Compute environments that use AWS Fargate
// resources must terminate all active jobs on that compute environment before
// deleting the compute environment. If this isn't done, the compute environment
// will end up in an invalid state.
func (c *Client) DeleteComputeEnvironment(ctx context.Context, params *DeleteComputeEnvironmentInput, optFns ...func(*Options)) (*DeleteComputeEnvironmentOutput, error) {
	if params == nil {
		params = &DeleteComputeEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteComputeEnvironment", params, optFns, addOperationDeleteComputeEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteComputeEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteComputeEnvironmentInput struct {

	// The name or Amazon Resource Name (ARN) of the compute environment to delete.
	//
	// This member is required.
	ComputeEnvironment *string
}

type DeleteComputeEnvironmentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteComputeEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteComputeEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteComputeEnvironment{}, middleware.After)
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
	if err = addOpDeleteComputeEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteComputeEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteComputeEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "DeleteComputeEnvironment",
	}
}
