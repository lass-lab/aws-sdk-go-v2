// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an log pattern to a LogPatternSet.
func (c *Client) CreateLogPattern(ctx context.Context, params *CreateLogPatternInput, optFns ...func(*Options)) (*CreateLogPatternOutput, error) {
	if params == nil {
		params = &CreateLogPatternInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLogPattern", params, optFns, addOperationCreateLogPatternMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLogPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogPatternInput struct {

	// The log pattern. The pattern must be DFA compatible. Patterns that utilize
	// forward lookahead or backreference constructions are not supported.
	//
	// This member is required.
	Pattern *string

	// The name of the log pattern.
	//
	// This member is required.
	PatternName *string

	// The name of the log pattern set.
	//
	// This member is required.
	PatternSetName *string

	// Rank of the log pattern. Must be a value between 1 and 1,000,000. The patterns
	// are sorted by rank, so we recommend that you set your highest priority patterns
	// with the lowest rank. A pattern of rank 1 will be the first to get matched to a
	// log line. A pattern of rank 1,000,000 will be last to get matched. When you
	// configure custom log patterns from the console, a Low severity pattern
	// translates to a 750,000 rank. A Medium severity pattern translates to a 500,000
	// rank. And a High severity pattern translates to a 250,000 rank. Rank values less
	// than 1 or greater than 1,000,000 are reserved for AWS-provided patterns.
	//
	// This member is required.
	Rank int32

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string
}

type CreateLogPatternOutput struct {

	// The successfully created log pattern.
	LogPattern *types.LogPattern

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLogPatternMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLogPattern{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLogPattern{}, middleware.After)
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
	if err = addOpCreateLogPatternValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLogPattern(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLogPattern(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "CreateLogPattern",
	}
}
