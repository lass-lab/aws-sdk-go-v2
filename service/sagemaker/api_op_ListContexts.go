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

// Lists the contexts in your account and their properties.
func (c *Client) ListContexts(ctx context.Context, params *ListContextsInput, optFns ...func(*Options)) (*ListContextsOutput, error) {
	if params == nil {
		params = &ListContextsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContexts", params, optFns, addOperationListContextsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContextsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContextsInput struct {

	// A filter that returns only contexts of the specified type.
	ContextType *string

	// A filter that returns only contexts created on or after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only contexts created on or before the specified time.
	CreatedBefore *time.Time

	// The maximum number of contexts to return in the response. The default value is
	// 10.
	MaxResults *int32

	// If the previous call to ListContexts didn't return the full set of contexts, the
	// call returns a token for getting the next set of contexts.
	NextToken *string

	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortContextsBy

	// The sort order. The default value is Descending.
	SortOrder types.SortOrder

	// A filter that returns only contexts with the specified source URI.
	SourceUri *string
}

type ListContextsOutput struct {

	// A list of contexts and their properties.
	ContextSummaries []types.ContextSummary

	// A token for getting the next set of contexts, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListContextsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListContexts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListContexts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContexts(options.Region), middleware.Before); err != nil {
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

// ListContextsAPIClient is a client that implements the ListContexts operation.
type ListContextsAPIClient interface {
	ListContexts(context.Context, *ListContextsInput, ...func(*Options)) (*ListContextsOutput, error)
}

var _ ListContextsAPIClient = (*Client)(nil)

// ListContextsPaginatorOptions is the paginator options for ListContexts
type ListContextsPaginatorOptions struct {
	// The maximum number of contexts to return in the response. The default value is
	// 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContextsPaginator is a paginator for ListContexts
type ListContextsPaginator struct {
	options   ListContextsPaginatorOptions
	client    ListContextsAPIClient
	params    *ListContextsInput
	nextToken *string
	firstPage bool
}

// NewListContextsPaginator returns a new ListContextsPaginator
func NewListContextsPaginator(client ListContextsAPIClient, params *ListContextsInput, optFns ...func(*ListContextsPaginatorOptions)) *ListContextsPaginator {
	options := ListContextsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListContextsInput{}
	}

	return &ListContextsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContextsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListContexts page.
func (p *ListContextsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContextsOutput, error) {
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

	result, err := p.client.ListContexts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContexts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListContexts",
	}
}
