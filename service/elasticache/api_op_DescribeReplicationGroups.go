// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns information about a particular replication group. If no identifier is
// specified, DescribeReplicationGroups returns information about all replication
// groups. This operation is valid for Redis only.
func (c *Client) DescribeReplicationGroups(ctx context.Context, params *DescribeReplicationGroupsInput, optFns ...func(*Options)) (*DescribeReplicationGroupsOutput, error) {
	if params == nil {
		params = &DescribeReplicationGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationGroups", params, optFns, addOperationDescribeReplicationGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeReplicationGroups operation.
type DescribeReplicationGroupsInput struct {

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32

	// The identifier for the replication group to be described. This parameter is not
	// case sensitive. If you do not specify this parameter, information about all
	// replication groups is returned.
	ReplicationGroupId *string
}

// Represents the output of a DescribeReplicationGroups operation.
type DescribeReplicationGroupsOutput struct {

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// A list of replication groups. Each item in the list contains detailed
	// information about one replication group.
	ReplicationGroups []types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReplicationGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReplicationGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReplicationGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationGroups(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationGroupsAPIClient is a client that implements the
// DescribeReplicationGroups operation.
type DescribeReplicationGroupsAPIClient interface {
	DescribeReplicationGroups(context.Context, *DescribeReplicationGroupsInput, ...func(*Options)) (*DescribeReplicationGroupsOutput, error)
}

var _ DescribeReplicationGroupsAPIClient = (*Client)(nil)

// DescribeReplicationGroupsPaginatorOptions is the paginator options for
// DescribeReplicationGroups
type DescribeReplicationGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationGroupsPaginator is a paginator for DescribeReplicationGroups
type DescribeReplicationGroupsPaginator struct {
	options   DescribeReplicationGroupsPaginatorOptions
	client    DescribeReplicationGroupsAPIClient
	params    *DescribeReplicationGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationGroupsPaginator returns a new
// DescribeReplicationGroupsPaginator
func NewDescribeReplicationGroupsPaginator(client DescribeReplicationGroupsAPIClient, params *DescribeReplicationGroupsInput, optFns ...func(*DescribeReplicationGroupsPaginatorOptions)) *DescribeReplicationGroupsPaginator {
	options := DescribeReplicationGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeReplicationGroupsInput{}
	}

	return &DescribeReplicationGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeReplicationGroups page.
func (p *DescribeReplicationGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeReplicationGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ReplicationGroupAvailableWaiterOptions are waiter options for
// ReplicationGroupAvailableWaiter
type ReplicationGroupAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ReplicationGroupAvailableWaiter will use default minimum delay of 15 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ReplicationGroupAvailableWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeReplicationGroupsInput, *DescribeReplicationGroupsOutput, error) (bool, error)
}

// ReplicationGroupAvailableWaiter defines the waiters for
// ReplicationGroupAvailable
type ReplicationGroupAvailableWaiter struct {
	client DescribeReplicationGroupsAPIClient

	options ReplicationGroupAvailableWaiterOptions
}

// NewReplicationGroupAvailableWaiter constructs a ReplicationGroupAvailableWaiter.
func NewReplicationGroupAvailableWaiter(client DescribeReplicationGroupsAPIClient, optFns ...func(*ReplicationGroupAvailableWaiterOptions)) *ReplicationGroupAvailableWaiter {
	options := ReplicationGroupAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = replicationGroupAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ReplicationGroupAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ReplicationGroupAvailable waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ReplicationGroupAvailableWaiter) Wait(ctx context.Context, params *DescribeReplicationGroupsInput, maxWaitDur time.Duration, optFns ...func(*ReplicationGroupAvailableWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeReplicationGroups(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ReplicationGroupAvailable waiter")
}

func replicationGroupAvailableStateRetryable(ctx context.Context, input *DescribeReplicationGroupsInput, output *DescribeReplicationGroupsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationGroups[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			if v != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationGroups[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		for _, v := range listOfValues {
			if v == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// ReplicationGroupDeletedWaiterOptions are waiter options for
// ReplicationGroupDeletedWaiter
type ReplicationGroupDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ReplicationGroupDeletedWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ReplicationGroupDeletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeReplicationGroupsInput, *DescribeReplicationGroupsOutput, error) (bool, error)
}

// ReplicationGroupDeletedWaiter defines the waiters for ReplicationGroupDeleted
type ReplicationGroupDeletedWaiter struct {
	client DescribeReplicationGroupsAPIClient

	options ReplicationGroupDeletedWaiterOptions
}

// NewReplicationGroupDeletedWaiter constructs a ReplicationGroupDeletedWaiter.
func NewReplicationGroupDeletedWaiter(client DescribeReplicationGroupsAPIClient, optFns ...func(*ReplicationGroupDeletedWaiterOptions)) *ReplicationGroupDeletedWaiter {
	options := ReplicationGroupDeletedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = replicationGroupDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ReplicationGroupDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ReplicationGroupDeleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ReplicationGroupDeletedWaiter) Wait(ctx context.Context, params *DescribeReplicationGroupsInput, maxWaitDur time.Duration, optFns ...func(*ReplicationGroupDeletedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeReplicationGroups(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ReplicationGroupDeleted waiter")
}

func replicationGroupDeletedStateRetryable(ctx context.Context, input *DescribeReplicationGroupsInput, output *DescribeReplicationGroupsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationGroups[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		var match = true
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			if v != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationGroups[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		for _, v := range listOfValues {
			if v == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err != nil {
		var errorType *types.ReplicationGroupNotFoundFault
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeReplicationGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeReplicationGroups",
	}
}
