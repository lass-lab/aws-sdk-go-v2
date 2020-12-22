// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a description of a model quality job definition.
func (c *Client) DescribeModelQualityJobDefinition(ctx context.Context, params *DescribeModelQualityJobDefinitionInput, optFns ...func(*Options)) (*DescribeModelQualityJobDefinitionOutput, error) {
	if params == nil {
		params = &DescribeModelQualityJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModelQualityJobDefinition", params, optFns, addOperationDescribeModelQualityJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelQualityJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelQualityJobDefinitionInput struct {

	// The name of the model quality job. The name must be unique within an AWS Region
	// in the AWS account.
	//
	// This member is required.
	JobDefinitionName *string
}

type DescribeModelQualityJobDefinitionOutput struct {

	// The time at which the model quality job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model quality job.
	//
	// This member is required.
	JobDefinitionArn *string

	// The name of the quality job definition. The name must be unique within an AWS
	// Region in the AWS account.
	//
	// This member is required.
	JobDefinitionName *string

	// Identifies the resources to deploy for a monitoring job.
	//
	// This member is required.
	JobResources *types.MonitoringResources

	// Configures the model quality job to run a specified Docker container image.
	//
	// This member is required.
	ModelQualityAppSpecification *types.ModelQualityAppSpecification

	// Inputs for the model quality job.
	//
	// This member is required.
	ModelQualityJobInput *types.ModelQualityJobInput

	// The output configuration for monitoring jobs.
	//
	// This member is required.
	ModelQualityJobOutputConfig *types.MonitoringOutputConfig

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	//
	// This member is required.
	RoleArn *string

	// The baseline configuration for a model quality job.
	ModelQualityBaselineConfig *types.ModelQualityBaselineConfig

	// Networking options for a model quality job.
	NetworkConfig *types.MonitoringNetworkConfig

	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition *types.MonitoringStoppingCondition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeModelQualityJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModelQualityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModelQualityJobDefinition{}, middleware.After)
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
	if err = addOpDescribeModelQualityJobDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModelQualityJobDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeModelQualityJobDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeModelQualityJobDefinition",
	}
}
