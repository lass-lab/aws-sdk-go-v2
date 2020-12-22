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
	"time"
)

// Lists the artifacts in your account and their properties.
func (c *Client) ListArtifacts(ctx context.Context, params *ListArtifactsInput, optFns ...func(*Options)) (*ListArtifactsOutput, error) {
	if params == nil {
		params = &ListArtifactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListArtifacts", params, optFns, addOperationListArtifactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListArtifactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListArtifactsInput struct {

	// A filter that returns only artifacts of the specified type.
	ArtifactType *string

	// A filter that returns only artifacts created on or after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only artifacts created on or before the specified time.
	CreatedBefore *time.Time

	// The maximum number of artifacts to return in the response. The default value is
	// 10.
	MaxResults *int32

	// If the previous call to ListArtifacts didn't return the full set of artifacts,
	// the call returns a token for getting the next set of artifacts.
	NextToken *string

	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortArtifactsBy

	// The sort order. The default value is Descending.
	SortOrder types.SortOrder

	// A filter that returns only artifacts with the specified source URI.
	SourceUri *string
}

type ListArtifactsOutput struct {

	// A list of artifacts and their properties.
	ArtifactSummaries []types.ArtifactSummary

	// A token for getting the next set of artifacts, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListArtifactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListArtifacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListArtifacts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListArtifacts(options.Region), middleware.Before); err != nil {
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

// ListArtifactsAPIClient is a client that implements the ListArtifacts operation.
type ListArtifactsAPIClient interface {
	ListArtifacts(context.Context, *ListArtifactsInput, ...func(*Options)) (*ListArtifactsOutput, error)
}

var _ ListArtifactsAPIClient = (*Client)(nil)

// ListArtifactsPaginatorOptions is the paginator options for ListArtifacts
type ListArtifactsPaginatorOptions struct {
	// The maximum number of artifacts to return in the response. The default value is
	// 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListArtifactsPaginator is a paginator for ListArtifacts
type ListArtifactsPaginator struct {
	options   ListArtifactsPaginatorOptions
	client    ListArtifactsAPIClient
	params    *ListArtifactsInput
	nextToken *string
	firstPage bool
}

// NewListArtifactsPaginator returns a new ListArtifactsPaginator
func NewListArtifactsPaginator(client ListArtifactsAPIClient, params *ListArtifactsInput, optFns ...func(*ListArtifactsPaginatorOptions)) *ListArtifactsPaginator {
	options := ListArtifactsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListArtifactsInput{}
	}

	return &ListArtifactsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListArtifactsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListArtifacts page.
func (p *ListArtifactsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListArtifactsOutput, error) {
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

	result, err := p.client.ListArtifacts(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListArtifacts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListArtifacts",
	}
}
