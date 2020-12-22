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

// Returns the optimization findings for an account. For example, it returns the
// number of Amazon EC2 instances in an account that are under-provisioned,
// over-provisioned, or optimized. It also returns the number of Auto Scaling
// groups in an account that are not optimized, or optimized.
func (c *Client) GetRecommendationSummaries(ctx context.Context, params *GetRecommendationSummariesInput, optFns ...func(*Options)) (*GetRecommendationSummariesOutput, error) {
	if params == nil {
		params = &GetRecommendationSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecommendationSummaries", params, optFns, addOperationGetRecommendationSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecommendationSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecommendationSummariesInput struct {

	// The IDs of the AWS accounts for which to return recommendation summaries. If
	// your account is the management account of an organization, use this parameter to
	// specify the member accounts for which you want to return recommendation
	// summaries. Only one account ID can be specified per request.
	AccountIds []string

	// The maximum number of recommendation summaries to return with a single request.
	// To retrieve the remaining results, make another request with the returned
	// NextToken value.
	MaxResults *int32

	// The token to advance to the next page of recommendation summaries.
	NextToken *string
}

type GetRecommendationSummariesOutput struct {

	// The token to use to advance to the next page of recommendation summaries. This
	// value is null when there are no more pages of recommendation summaries to
	// return.
	NextToken *string

	// An array of objects that summarize a recommendation.
	RecommendationSummaries []types.RecommendationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRecommendationSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRecommendationSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRecommendationSummaries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendationSummaries(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRecommendationSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetRecommendationSummaries",
	}
}
