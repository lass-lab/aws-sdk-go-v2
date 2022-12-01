// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the PII entity detection jobs that you have submitted.
func (c *Client) ListPiiEntitiesDetectionJobs(ctx context.Context, params *ListPiiEntitiesDetectionJobsInput, optFns ...func(*Options)) (*ListPiiEntitiesDetectionJobsOutput, error) {
	if params == nil {
		params = &ListPiiEntitiesDetectionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPiiEntitiesDetectionJobs", params, optFns, c.addOperationListPiiEntitiesDetectionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPiiEntitiesDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPiiEntitiesDetectionJobsInput struct {

	// Filters the jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.PiiEntitiesDetectionJobFilter

	// The maximum number of results to return in each page.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPiiEntitiesDetectionJobsOutput struct {

	// Identifies the next page of results to return.
	NextToken *string

	// A list containing the properties of each job that is returned.
	PiiEntitiesDetectionJobPropertiesList []types.PiiEntitiesDetectionJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPiiEntitiesDetectionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPiiEntitiesDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPiiEntitiesDetectionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPiiEntitiesDetectionJobs(options.Region), middleware.Before); err != nil {
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

// ListPiiEntitiesDetectionJobsAPIClient is a client that implements the
// ListPiiEntitiesDetectionJobs operation.
type ListPiiEntitiesDetectionJobsAPIClient interface {
	ListPiiEntitiesDetectionJobs(context.Context, *ListPiiEntitiesDetectionJobsInput, ...func(*Options)) (*ListPiiEntitiesDetectionJobsOutput, error)
}

var _ ListPiiEntitiesDetectionJobsAPIClient = (*Client)(nil)

// ListPiiEntitiesDetectionJobsPaginatorOptions is the paginator options for
// ListPiiEntitiesDetectionJobs
type ListPiiEntitiesDetectionJobsPaginatorOptions struct {
	// The maximum number of results to return in each page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPiiEntitiesDetectionJobsPaginator is a paginator for
// ListPiiEntitiesDetectionJobs
type ListPiiEntitiesDetectionJobsPaginator struct {
	options   ListPiiEntitiesDetectionJobsPaginatorOptions
	client    ListPiiEntitiesDetectionJobsAPIClient
	params    *ListPiiEntitiesDetectionJobsInput
	nextToken *string
	firstPage bool
}

// NewListPiiEntitiesDetectionJobsPaginator returns a new
// ListPiiEntitiesDetectionJobsPaginator
func NewListPiiEntitiesDetectionJobsPaginator(client ListPiiEntitiesDetectionJobsAPIClient, params *ListPiiEntitiesDetectionJobsInput, optFns ...func(*ListPiiEntitiesDetectionJobsPaginatorOptions)) *ListPiiEntitiesDetectionJobsPaginator {
	if params == nil {
		params = &ListPiiEntitiesDetectionJobsInput{}
	}

	options := ListPiiEntitiesDetectionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPiiEntitiesDetectionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPiiEntitiesDetectionJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPiiEntitiesDetectionJobs page.
func (p *ListPiiEntitiesDetectionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPiiEntitiesDetectionJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListPiiEntitiesDetectionJobs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPiiEntitiesDetectionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListPiiEntitiesDetectionJobs",
	}
}
