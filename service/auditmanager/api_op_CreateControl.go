// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new custom control in AWS Audit Manager.
func (c *Client) CreateControl(ctx context.Context, params *CreateControlInput, optFns ...func(*Options)) (*CreateControlOutput, error) {
	if params == nil {
		params = &CreateControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateControl", params, optFns, addOperationCreateControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateControlInput struct {

	// The data source that determines from where AWS Audit Manager collects evidence
	// for the control.
	//
	// This member is required.
	ControlMappingSources []types.CreateControlMappingSource

	// The name of the control.
	//
	// This member is required.
	Name *string

	// The recommended actions to carry out if the control is not fulfilled.
	ActionPlanInstructions *string

	// The title of the action plan for remediating the control.
	ActionPlanTitle *string

	// The description of the control.
	Description *string

	// The tags associated with the control.
	Tags map[string]string

	// The steps to follow to determine if the control has been satisfied.
	TestingInformation *string
}

type CreateControlOutput struct {

	// The new control returned by the CreateControl API.
	Control *types.Control

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateControl{}, middleware.After)
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
	if err = addOpCreateControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "CreateControl",
	}
}
