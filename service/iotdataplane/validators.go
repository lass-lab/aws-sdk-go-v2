// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDeleteThingShadow struct {
}

func (*validateOpDeleteThingShadow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteThingShadow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteThingShadowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteThingShadowInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetThingShadow struct {
}

func (*validateOpGetThingShadow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetThingShadow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetThingShadowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetThingShadowInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListNamedShadowsForThing struct {
}

func (*validateOpListNamedShadowsForThing) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListNamedShadowsForThing) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListNamedShadowsForThingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListNamedShadowsForThingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPublish struct {
}

func (*validateOpPublish) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPublish) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PublishInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPublishInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateThingShadow struct {
}

func (*validateOpUpdateThingShadow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateThingShadow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateThingShadowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateThingShadowInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteThingShadowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteThingShadow{}, middleware.After)
}

func addOpGetThingShadowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetThingShadow{}, middleware.After)
}

func addOpListNamedShadowsForThingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListNamedShadowsForThing{}, middleware.After)
}

func addOpPublishValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPublish{}, middleware.After)
}

func addOpUpdateThingShadowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateThingShadow{}, middleware.After)
}

func validateOpDeleteThingShadowInput(v *DeleteThingShadowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteThingShadowInput"}
	if v.ThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetThingShadowInput(v *GetThingShadowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetThingShadowInput"}
	if v.ThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListNamedShadowsForThingInput(v *ListNamedShadowsForThingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListNamedShadowsForThingInput"}
	if v.ThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPublishInput(v *PublishInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PublishInput"}
	if v.Topic == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Topic"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateThingShadowInput(v *UpdateThingShadowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateThingShadowInput"}
	if v.ThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThingName"))
	}
	if v.Payload == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Payload"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
