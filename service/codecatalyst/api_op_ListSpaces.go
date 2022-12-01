// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of spaces.
func (c *Client) ListSpaces(ctx context.Context, params *ListSpacesInput, optFns ...func(*Options)) (*ListSpacesOutput, error) {
	if params == nil {
		params = &ListSpacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSpaces", params, optFns, c.addOperationListSpacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSpacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSpacesInput struct {

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSpacesOutput struct {

	// Information about the space.
	Items []types.SpaceSummary

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSpacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSpaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSpaces{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addBearerAuthSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSpaces(options.Region), middleware.Before); err != nil {
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

// ListSpacesAPIClient is a client that implements the ListSpaces operation.
type ListSpacesAPIClient interface {
	ListSpaces(context.Context, *ListSpacesInput, ...func(*Options)) (*ListSpacesOutput, error)
}

var _ ListSpacesAPIClient = (*Client)(nil)

// ListSpacesPaginatorOptions is the paginator options for ListSpaces
type ListSpacesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSpacesPaginator is a paginator for ListSpaces
type ListSpacesPaginator struct {
	options   ListSpacesPaginatorOptions
	client    ListSpacesAPIClient
	params    *ListSpacesInput
	nextToken *string
	firstPage bool
}

// NewListSpacesPaginator returns a new ListSpacesPaginator
func NewListSpacesPaginator(client ListSpacesAPIClient, params *ListSpacesInput, optFns ...func(*ListSpacesPaginatorOptions)) *ListSpacesPaginator {
	if params == nil {
		params = &ListSpacesInput{}
	}

	options := ListSpacesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSpacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSpacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSpaces page.
func (p *ListSpacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSpacesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListSpaces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSpaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSpaces",
	}
}
