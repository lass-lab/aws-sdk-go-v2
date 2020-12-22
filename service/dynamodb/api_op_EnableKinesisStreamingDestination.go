// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts table data replication to the specified Kinesis data stream at a
// timestamp chosen during the enable workflow. If this operation doesn't return
// results immediately, use DescribeKinesisStreamingDestination to check if
// streaming to the Kinesis data stream is ACTIVE.
func (c *Client) EnableKinesisStreamingDestination(ctx context.Context, params *EnableKinesisStreamingDestinationInput, optFns ...func(*Options)) (*EnableKinesisStreamingDestinationOutput, error) {
	if params == nil {
		params = &EnableKinesisStreamingDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableKinesisStreamingDestination", params, optFns, addOperationEnableKinesisStreamingDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableKinesisStreamingDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableKinesisStreamingDestinationInput struct {

	// The ARN for a Kinesis data stream.
	//
	// This member is required.
	StreamArn *string

	// The name of the DynamoDB table.
	//
	// This member is required.
	TableName *string
}

type EnableKinesisStreamingDestinationOutput struct {

	// The current status of the replication.
	DestinationStatus types.DestinationStatus

	// The ARN for the specific Kinesis data stream.
	StreamArn *string

	// The name of the table being modified.
	TableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableKinesisStreamingDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpEnableKinesisStreamingDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpEnableKinesisStreamingDestination{}, middleware.After)
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
	if err = addOpEnableKinesisStreamingDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableKinesisStreamingDestination(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opEnableKinesisStreamingDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "EnableKinesisStreamingDestination",
	}
}
