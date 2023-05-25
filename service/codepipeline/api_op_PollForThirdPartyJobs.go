// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Determines whether there are any third party jobs for a job worker to act on.
// Used for partner actions only. When this API is called, CodePipeline returns
// temporary credentials for the S3 bucket used to store artifacts for the
// pipeline, if the action requires access to that S3 bucket for input or output
// artifacts.
func (c *Client) PollForThirdPartyJobs(ctx context.Context, params *PollForThirdPartyJobsInput, optFns ...func(*Options)) (*PollForThirdPartyJobsOutput, error) {
	if params == nil {
		params = &PollForThirdPartyJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PollForThirdPartyJobs", params, optFns, c.addOperationPollForThirdPartyJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PollForThirdPartyJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PollForThirdPartyJobs action.
type PollForThirdPartyJobsInput struct {

	// Represents information about an action type.
	//
	// This member is required.
	ActionTypeId *types.ActionTypeId

	// The maximum number of jobs to return in a poll for jobs call.
	MaxBatchSize *int32

	noSmithyDocumentSerde
}

// Represents the output of a PollForThirdPartyJobs action.
type PollForThirdPartyJobsOutput struct {

	// Information about the jobs to take action on.
	Jobs []types.ThirdPartyJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPollForThirdPartyJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPollForThirdPartyJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPollForThirdPartyJobs{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpPollForThirdPartyJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPollForThirdPartyJobs(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opPollForThirdPartyJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PollForThirdPartyJobs",
	}
}
