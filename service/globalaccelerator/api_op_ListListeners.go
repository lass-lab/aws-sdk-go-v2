// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the listeners for an accelerator.
func (c *Client) ListListeners(ctx context.Context, params *ListListenersInput, optFns ...func(*Options)) (*ListListenersOutput, error) {
	if params == nil {
		params = &ListListenersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListListeners", params, optFns, addOperationListListenersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListListenersInput struct {

	// The Amazon Resource Name (ARN) of the accelerator for which you want to list
	// listener objects.
	//
	// This member is required.
	AcceleratorArn *string

	// The number of listener objects that you want to return with this call. The
	// default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
}

type ListListenersOutput struct {

	// The list of listeners for an accelerator.
	Listeners []types.Listener

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListListenersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListListeners{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListListeners{}, middleware.After)
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
	if err = addOpListListenersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListListeners(options.Region), middleware.Before); err != nil {
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

// ListListenersAPIClient is a client that implements the ListListeners operation.
type ListListenersAPIClient interface {
	ListListeners(context.Context, *ListListenersInput, ...func(*Options)) (*ListListenersOutput, error)
}

var _ ListListenersAPIClient = (*Client)(nil)

// ListListenersPaginatorOptions is the paginator options for ListListeners
type ListListenersPaginatorOptions struct {
	// The number of listener objects that you want to return with this call. The
	// default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListListenersPaginator is a paginator for ListListeners
type ListListenersPaginator struct {
	options   ListListenersPaginatorOptions
	client    ListListenersAPIClient
	params    *ListListenersInput
	nextToken *string
	firstPage bool
}

// NewListListenersPaginator returns a new ListListenersPaginator
func NewListListenersPaginator(client ListListenersAPIClient, params *ListListenersInput, optFns ...func(*ListListenersPaginatorOptions)) *ListListenersPaginator {
	options := ListListenersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListListenersInput{}
	}

	return &ListListenersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListListenersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListListeners page.
func (p *ListListenersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListListenersOutput, error) {
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

	result, err := p.client.ListListeners(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListListeners(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListListeners",
	}
}
