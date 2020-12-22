// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits an AWS Batch job from a job definition. Parameters specified during
// SubmitJob override parameters defined in the job definition. Jobs run on Fargate
// resources don't run for more than 14 days. After 14 days, the Fargate resources
// might no longer be available and the job is terminated.
func (c *Client) SubmitJob(ctx context.Context, params *SubmitJobInput, optFns ...func(*Options)) (*SubmitJobOutput, error) {
	if params == nil {
		params = &SubmitJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubmitJob", params, optFns, addOperationSubmitJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubmitJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitJobInput struct {

	// The job definition used by this job. This value can be one of name,
	// name:revision, or the Amazon Resource Name (ARN) for the job definition. If name
	// is specified without a revision then the latest active revision is used.
	//
	// This member is required.
	JobDefinition *string

	// The name of the job. The first character must be alphanumeric, and up to 128
	// letters (uppercase and lowercase), numbers, hyphens, and underscores are
	// allowed.
	//
	// This member is required.
	JobName *string

	// The job queue into which the job is submitted. You can specify either the name
	// or the Amazon Resource Name (ARN) of the queue.
	//
	// This member is required.
	JobQueue *string

	// The array properties for the submitted job, such as the size of the array. The
	// array size can be between 2 and 10,000. If you specify array properties for a
	// job, it becomes an array job. For more information, see Array Jobs
	// (https://docs.aws.amazon.com/batch/latest/userguide/array_jobs.html) in the AWS
	// Batch User Guide.
	ArrayProperties *types.ArrayProperties

	// A list of container overrides in JSON format that specify the name of a
	// container in the specified job definition and the overrides it should receive.
	// You can override the default command for a container (that is specified in the
	// job definition or the Docker image) with a command override. You can also
	// override existing environment variables (that are specified in the job
	// definition or Docker image) on a container or add new environment variables to
	// it with an environment override.
	ContainerOverrides *types.ContainerOverrides

	// A list of dependencies for the job. A job can depend upon a maximum of 20 jobs.
	// You can specify a SEQUENTIAL type dependency without specifying a job ID for
	// array jobs so that each child array job completes sequentially, starting at
	// index 0. You can also specify an N_TO_N type dependency with a job ID for array
	// jobs. In that case, each index child of this job must wait for the corresponding
	// index child of each dependency to complete before it can begin.
	DependsOn []types.JobDependency

	// A list of node overrides in JSON format that specify the node range to target
	// and the container overrides for that node range. This parameter isn't applicable
	// to jobs running on Fargate resources; use containerOverrides instead.
	NodeOverrides *types.NodeOverrides

	// Additional parameters passed to the job that replace parameter substitution
	// placeholders that are set in the job definition. Parameters are specified as a
	// key and value pair mapping. Parameters in a SubmitJob request override any
	// corresponding parameter defaults from the job definition.
	Parameters map[string]string

	// Specifies whether to propagate the tags from the job or job definition to the
	// corresponding Amazon ECS task. If no value is specified, the tags aren't
	// propagated. Tags can only be propagated to the tasks during task creation. For
	// tags with the same name, job tags are given priority over job definitions tags.
	// If the total number of combined tags from the job and job definition is over 50,
	// the job is moved to the FAILED state. When specified, this overrides the tag
	// propagation setting in the job definition.
	PropagateTags bool

	// The retry strategy to use for failed jobs from this SubmitJob operation. When a
	// retry strategy is specified here, it overrides the retry strategy defined in the
	// job definition.
	RetryStrategy *types.RetryStrategy

	// The tags that you apply to the job request to help you categorize and organize
	// your resources. Each tag consists of a key and an optional value. For more
	// information, see Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in AWS General
	// Reference.
	Tags map[string]string

	// The timeout configuration for this SubmitJob operation. You can specify a
	// timeout duration after which AWS Batch terminates your jobs if they haven't
	// finished. If a job is terminated due to a timeout, it isn't retried. The minimum
	// value for the timeout is 60 seconds. This configuration overrides any timeout
	// configuration specified in the job definition. For array jobs, child jobs have
	// the same timeout configuration as the parent job. For more information, see Job
	// Timeouts
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/job_timeouts.html)
	// in the Amazon Elastic Container Service Developer Guide.
	Timeout *types.JobTimeout
}

type SubmitJobOutput struct {

	// The unique identifier for the job.
	//
	// This member is required.
	JobId *string

	// The name of the job.
	//
	// This member is required.
	JobName *string

	// The Amazon Resource Name (ARN) for the job.
	JobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSubmitJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSubmitJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSubmitJob{}, middleware.After)
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
	if err = addOpSubmitJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubmitJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "SubmitJob",
	}
}
