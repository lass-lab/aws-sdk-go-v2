// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a pipeline using a JSON pipeline definition.
func (c *Client) CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) {
	if params == nil {
		params = &CreatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePipeline", params, optFns, addOperationCreatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePipelineInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time.
	//
	// This member is required.
	ClientRequestToken *string

	// The JSON pipeline definition of the pipeline.
	//
	// This member is required.
	PipelineDefinition *string

	// The name of the pipeline.
	//
	// This member is required.
	PipelineName *string

	// The Amazon Resource Name (ARN) of the role used by the pipeline to access and
	// create resources.
	//
	// This member is required.
	RoleArn *string

	// A description of the pipeline.
	PipelineDescription *string

	// The display name of the pipeline.
	PipelineDisplayName *string

	// A list of tags to apply to the created pipeline.
	Tags []types.Tag
}

type CreatePipelineOutput struct {

	// The Amazon Resource Name (ARN) of the created pipeline.
	PipelineArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePipeline{}, middleware.After)
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
	if err = addIdempotencyToken_opCreatePipelineMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePipeline(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePipeline struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePipeline) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePipelineInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePipelineMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePipeline{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreatePipeline",
	}
}
