// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// This operation returns information about a vault, including the vault's Amazon
// Resource Name (ARN), the date the vault was created, the number of archives it
// contains, and the total size of all the archives in the vault. The number of
// archives and their total size are as of the last inventory generation. This
// means that if you add or remove an archive from a vault, and then immediately
// use Describe Vault, the change in contents will not be immediately reflected. If
// you want to retrieve the latest inventory of the vault, use InitiateJob. Amazon
// S3 Glacier generates vault inventories approximately daily. For more
// information, see Downloading a Vault Inventory in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-inventory.html). An
// AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions by
// default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access
// Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For conceptual information and underlying REST API, see Retrieving Vault
// Metadata in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/retrieving-vault-info.html)
// and Describe Vault
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-get.html) in the
// Amazon Glacier Developer Guide.
func (c *Client) DescribeVault(ctx context.Context, params *DescribeVaultInput, optFns ...func(*Options)) (*DescribeVaultOutput, error) {
	if params == nil {
		params = &DescribeVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVault", params, optFns, addOperationDescribeVaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving metadata for a specific vault in Amazon Glacier.
type DescribeVaultInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

// Contains the Amazon S3 Glacier response to your request.
type DescribeVaultOutput struct {

	// The Universal Coordinated Time (UTC) date when the vault was created. This value
	// should be a string in the ISO 8601 date format, for example
	// 2012-03-20T17:03:43.221Z.
	CreationDate *string

	// The Universal Coordinated Time (UTC) date when Amazon S3 Glacier completed the
	// last vault inventory. This value should be a string in the ISO 8601 date format,
	// for example 2012-03-20T17:03:43.221Z.
	LastInventoryDate *string

	// The number of archives in the vault as of the last inventory date. This field
	// will return null if an inventory has not yet run on the vault, for example if
	// you just created the vault.
	NumberOfArchives int64

	// Total size, in bytes, of the archives in the vault as of the last inventory
	// date. This field will return null if an inventory has not yet run on the vault,
	// for example if you just created the vault.
	SizeInBytes int64

	// The Amazon Resource Name (ARN) of the vault.
	VaultARN *string

	// The name of the vault.
	VaultName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVault{}, middleware.After)
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
	if err = addOpDescribeVaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVault(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeVaultAPIClient is a client that implements the DescribeVault operation.
type DescribeVaultAPIClient interface {
	DescribeVault(context.Context, *DescribeVaultInput, ...func(*Options)) (*DescribeVaultOutput, error)
}

var _ DescribeVaultAPIClient = (*Client)(nil)

// VaultExistsWaiterOptions are waiter options for VaultExistsWaiter
type VaultExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VaultExistsWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, VaultExistsWaiter will use default max delay of 120 seconds. Note that
	// MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeVaultInput, *DescribeVaultOutput, error) (bool, error)
}

// VaultExistsWaiter defines the waiters for VaultExists
type VaultExistsWaiter struct {
	client DescribeVaultAPIClient

	options VaultExistsWaiterOptions
}

// NewVaultExistsWaiter constructs a VaultExistsWaiter.
func NewVaultExistsWaiter(client DescribeVaultAPIClient, optFns ...func(*VaultExistsWaiterOptions)) *VaultExistsWaiter {
	options := VaultExistsWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = vaultExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VaultExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VaultExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *VaultExistsWaiter) Wait(ctx context.Context, params *DescribeVaultInput, maxWaitDur time.Duration, optFns ...func(*VaultExistsWaiterOptions)) error {
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

		out, err := w.client.DescribeVault(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for VaultExists waiter")
}

func vaultExistsStateRetryable(ctx context.Context, input *DescribeVaultInput, output *DescribeVaultOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// VaultNotExistsWaiterOptions are waiter options for VaultNotExistsWaiter
type VaultNotExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VaultNotExistsWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, VaultNotExistsWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeVaultInput, *DescribeVaultOutput, error) (bool, error)
}

// VaultNotExistsWaiter defines the waiters for VaultNotExists
type VaultNotExistsWaiter struct {
	client DescribeVaultAPIClient

	options VaultNotExistsWaiterOptions
}

// NewVaultNotExistsWaiter constructs a VaultNotExistsWaiter.
func NewVaultNotExistsWaiter(client DescribeVaultAPIClient, optFns ...func(*VaultNotExistsWaiterOptions)) *VaultNotExistsWaiter {
	options := VaultNotExistsWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = vaultNotExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VaultNotExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VaultNotExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *VaultNotExistsWaiter) Wait(ctx context.Context, params *DescribeVaultInput, maxWaitDur time.Duration, optFns ...func(*VaultNotExistsWaiterOptions)) error {
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

		out, err := w.client.DescribeVault(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for VaultNotExists waiter")
}

func vaultNotExistsStateRetryable(ctx context.Context, input *DescribeVaultInput, output *DescribeVaultOutput, err error) (bool, error) {

	if err == nil {
		return true, nil
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeVault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "DescribeVault",
	}
}
