// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the tags that are attached to the specified IAM customer managed policy.
// The returned list of tags is sorted by tag key. For more information about
// tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
// in the IAM User Guide.
func (c *Client) ListPolicyTags(ctx context.Context, params *ListPolicyTagsInput, optFns ...func(*Options)) (*ListPolicyTagsOutput, error) {
	if params == nil {
		params = &ListPolicyTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicyTags", params, optFns, c.addOperationListPolicyTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPolicyTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPolicyTagsInput struct {

	// The ARN of the IAM customer managed policy whose tags you want to see. This
	// parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex) )
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	PolicyArn *string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListPolicyTagsOutput struct {

	// The list of tags that are currently attached to the IAM customer managed
	// policy. Each tag consists of a key name and an associated value. If no tags are
	// attached to the specified resource, the response contains an empty list.
	//
	// This member is required.
	Tags []types.Tag

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPolicyTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListPolicyTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListPolicyTags{}, middleware.After)
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListPolicyTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyTags(options.Region), middleware.Before); err != nil {
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

// ListPolicyTagsAPIClient is a client that implements the ListPolicyTags
// operation.
type ListPolicyTagsAPIClient interface {
	ListPolicyTags(context.Context, *ListPolicyTagsInput, ...func(*Options)) (*ListPolicyTagsOutput, error)
}

var _ ListPolicyTagsAPIClient = (*Client)(nil)

// ListPolicyTagsPaginatorOptions is the paginator options for ListPolicyTags
type ListPolicyTagsPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPolicyTagsPaginator is a paginator for ListPolicyTags
type ListPolicyTagsPaginator struct {
	options   ListPolicyTagsPaginatorOptions
	client    ListPolicyTagsAPIClient
	params    *ListPolicyTagsInput
	nextToken *string
	firstPage bool
}

// NewListPolicyTagsPaginator returns a new ListPolicyTagsPaginator
func NewListPolicyTagsPaginator(client ListPolicyTagsAPIClient, params *ListPolicyTagsInput, optFns ...func(*ListPolicyTagsPaginatorOptions)) *ListPolicyTagsPaginator {
	if params == nil {
		params = &ListPolicyTagsInput{}
	}

	options := ListPolicyTagsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPolicyTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPolicyTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPolicyTags page.
func (p *ListPolicyTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPolicyTagsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListPolicyTags(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPolicyTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListPolicyTags",
	}
}
