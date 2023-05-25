// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a listing of all the webhooks in this Amazon Web Services Region for this
// account. The output lists all webhooks and includes the webhook URL and ARN and
// the configuration for each webhook.
func (c *Client) ListWebhooks(ctx context.Context, params *ListWebhooksInput, optFns ...func(*Options)) (*ListWebhooksOutput, error) {
	if params == nil {
		params = &ListWebhooksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWebhooks", params, optFns, c.addOperationListWebhooksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWebhooksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWebhooksInput struct {

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token that was returned from the previous ListWebhooks call, which can be
	// used to return the next set of webhooks in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWebhooksOutput struct {

	// If the amount of returned information is significantly large, an identifier is
	// also returned and can be used in a subsequent ListWebhooks call to return the
	// next set of webhooks in the list.
	NextToken *string

	// The JSON detail returned for each webhook in the list output for the
	// ListWebhooks call.
	Webhooks []types.ListWebhookItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWebhooksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWebhooks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWebhooks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWebhooks(options.Region), middleware.Before); err != nil {
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

// ListWebhooksAPIClient is a client that implements the ListWebhooks operation.
type ListWebhooksAPIClient interface {
	ListWebhooks(context.Context, *ListWebhooksInput, ...func(*Options)) (*ListWebhooksOutput, error)
}

var _ ListWebhooksAPIClient = (*Client)(nil)

// ListWebhooksPaginatorOptions is the paginator options for ListWebhooks
type ListWebhooksPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWebhooksPaginator is a paginator for ListWebhooks
type ListWebhooksPaginator struct {
	options   ListWebhooksPaginatorOptions
	client    ListWebhooksAPIClient
	params    *ListWebhooksInput
	nextToken *string
	firstPage bool
}

// NewListWebhooksPaginator returns a new ListWebhooksPaginator
func NewListWebhooksPaginator(client ListWebhooksAPIClient, params *ListWebhooksInput, optFns ...func(*ListWebhooksPaginatorOptions)) *ListWebhooksPaginator {
	if params == nil {
		params = &ListWebhooksInput{}
	}

	options := ListWebhooksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWebhooksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWebhooksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWebhooks page.
func (p *ListWebhooksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWebhooksOutput, error) {
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

	result, err := p.client.ListWebhooks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWebhooks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "ListWebhooks",
	}
}
