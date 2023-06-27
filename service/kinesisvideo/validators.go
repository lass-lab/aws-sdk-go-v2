// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateSignalingChannel struct {
}

func (*validateOpCreateSignalingChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSignalingChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSignalingChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSignalingChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateStream struct {
}

func (*validateOpCreateStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSignalingChannel struct {
}

func (*validateOpDeleteSignalingChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSignalingChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSignalingChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSignalingChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteStream struct {
}

func (*validateOpDeleteStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDataEndpoint struct {
}

func (*validateOpGetDataEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDataEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDataEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDataEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSignalingChannelEndpoint struct {
}

func (*validateOpGetSignalingChannelEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSignalingChannelEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSignalingChannelEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSignalingChannelEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEdgeAgentConfigurations struct {
}

func (*validateOpListEdgeAgentConfigurations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEdgeAgentConfigurations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEdgeAgentConfigurationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEdgeAgentConfigurationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartEdgeConfigurationUpdate struct {
}

func (*validateOpStartEdgeConfigurationUpdate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartEdgeConfigurationUpdate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartEdgeConfigurationUpdateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartEdgeConfigurationUpdateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagStream struct {
}

func (*validateOpTagStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagStream struct {
}

func (*validateOpUntagStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDataRetention struct {
}

func (*validateOpUpdateDataRetention) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDataRetention) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDataRetentionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDataRetentionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateImageGenerationConfiguration struct {
}

func (*validateOpUpdateImageGenerationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateImageGenerationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateImageGenerationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateImageGenerationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateMediaStorageConfiguration struct {
}

func (*validateOpUpdateMediaStorageConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMediaStorageConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMediaStorageConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMediaStorageConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateNotificationConfiguration struct {
}

func (*validateOpUpdateNotificationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateNotificationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateNotificationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateNotificationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSignalingChannel struct {
}

func (*validateOpUpdateSignalingChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSignalingChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSignalingChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSignalingChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateStream struct {
}

func (*validateOpUpdateStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateSignalingChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSignalingChannel{}, middleware.After)
}

func addOpCreateStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateStream{}, middleware.After)
}

func addOpDeleteSignalingChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSignalingChannel{}, middleware.After)
}

func addOpDeleteStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteStream{}, middleware.After)
}

func addOpGetDataEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDataEndpoint{}, middleware.After)
}

func addOpGetSignalingChannelEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSignalingChannelEndpoint{}, middleware.After)
}

func addOpListEdgeAgentConfigurationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEdgeAgentConfigurations{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartEdgeConfigurationUpdateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartEdgeConfigurationUpdate{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpTagStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagStream{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUntagStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagStream{}, middleware.After)
}

func addOpUpdateDataRetentionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDataRetention{}, middleware.After)
}

func addOpUpdateImageGenerationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateImageGenerationConfiguration{}, middleware.After)
}

func addOpUpdateMediaStorageConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMediaStorageConfiguration{}, middleware.After)
}

func addOpUpdateNotificationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateNotificationConfiguration{}, middleware.After)
}

func addOpUpdateSignalingChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSignalingChannel{}, middleware.After)
}

func addOpUpdateStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateStream{}, middleware.After)
}

func validateEdgeConfig(v *types.EdgeConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EdgeConfig"}
	if v.HubDeviceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HubDeviceArn"))
	}
	if v.RecorderConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RecorderConfig"))
	} else if v.RecorderConfig != nil {
		if err := validateRecorderConfig(v.RecorderConfig); err != nil {
			invalidParams.AddNested("RecorderConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.UploaderConfig != nil {
		if err := validateUploaderConfig(v.UploaderConfig); err != nil {
			invalidParams.AddNested("UploaderConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateImageGenerationConfiguration(v *types.ImageGenerationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImageGenerationConfiguration"}
	if len(v.Status) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Status"))
	}
	if len(v.ImageSelectorType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ImageSelectorType"))
	}
	if v.DestinationConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationConfig"))
	} else if v.DestinationConfig != nil {
		if err := validateImageGenerationDestinationConfig(v.DestinationConfig); err != nil {
			invalidParams.AddNested("DestinationConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.SamplingInterval == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SamplingInterval"))
	}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateImageGenerationDestinationConfig(v *types.ImageGenerationDestinationConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImageGenerationDestinationConfig"}
	if v.Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Uri"))
	}
	if v.DestinationRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationRegion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMediaSourceConfig(v *types.MediaSourceConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MediaSourceConfig"}
	if v.MediaUriSecretArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MediaUriSecretArn"))
	}
	if len(v.MediaUriType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MediaUriType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMediaStorageConfiguration(v *types.MediaStorageConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MediaStorageConfiguration"}
	if len(v.Status) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Status"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNotificationConfiguration(v *types.NotificationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NotificationConfiguration"}
	if len(v.Status) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Status"))
	}
	if v.DestinationConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationConfig"))
	} else if v.DestinationConfig != nil {
		if err := validateNotificationDestinationConfig(v.DestinationConfig); err != nil {
			invalidParams.AddNested("DestinationConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNotificationDestinationConfig(v *types.NotificationDestinationConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NotificationDestinationConfig"}
	if v.Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRecorderConfig(v *types.RecorderConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RecorderConfig"}
	if v.MediaSourceConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MediaSourceConfig"))
	} else if v.MediaSourceConfig != nil {
		if err := validateMediaSourceConfig(v.MediaSourceConfig); err != nil {
			invalidParams.AddNested("MediaSourceConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.ScheduleConfig != nil {
		if err := validateScheduleConfig(v.ScheduleConfig); err != nil {
			invalidParams.AddNested("ScheduleConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScheduleConfig(v *types.ScheduleConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScheduleConfig"}
	if v.ScheduleExpression == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduleExpression"))
	}
	if v.DurationInSeconds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DurationInSeconds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagOnCreateList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagOnCreateList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUploaderConfig(v *types.UploaderConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UploaderConfig"}
	if v.ScheduleConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduleConfig"))
	} else if v.ScheduleConfig != nil {
		if err := validateScheduleConfig(v.ScheduleConfig); err != nil {
			invalidParams.AddNested("ScheduleConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSignalingChannelInput(v *CreateSignalingChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSignalingChannelInput"}
	if v.ChannelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelName"))
	}
	if v.Tags != nil {
		if err := validateTagOnCreateList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateStreamInput(v *CreateStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateStreamInput"}
	if v.StreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSignalingChannelInput(v *DeleteSignalingChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSignalingChannelInput"}
	if v.ChannelARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteStreamInput(v *DeleteStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteStreamInput"}
	if v.StreamARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDataEndpointInput(v *GetDataEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDataEndpointInput"}
	if len(v.APIName) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("APIName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSignalingChannelEndpointInput(v *GetSignalingChannelEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSignalingChannelEndpointInput"}
	if v.ChannelARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEdgeAgentConfigurationsInput(v *ListEdgeAgentConfigurationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEdgeAgentConfigurationsInput"}
	if v.HubDeviceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HubDeviceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartEdgeConfigurationUpdateInput(v *StartEdgeConfigurationUpdateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartEdgeConfigurationUpdateInput"}
	if v.EdgeConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EdgeConfig"))
	} else if v.EdgeConfig != nil {
		if err := validateEdgeConfig(v.EdgeConfig); err != nil {
			invalidParams.AddNested("EdgeConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagStreamInput(v *TagStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagStreamInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeyList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeyList"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagStreamInput(v *UntagStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagStreamInput"}
	if v.TagKeyList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeyList"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateDataRetentionInput(v *UpdateDataRetentionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDataRetentionInput"}
	if v.CurrentVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CurrentVersion"))
	}
	if len(v.Operation) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Operation"))
	}
	if v.DataRetentionChangeInHours == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataRetentionChangeInHours"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateImageGenerationConfigurationInput(v *UpdateImageGenerationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateImageGenerationConfigurationInput"}
	if v.ImageGenerationConfiguration != nil {
		if err := validateImageGenerationConfiguration(v.ImageGenerationConfiguration); err != nil {
			invalidParams.AddNested("ImageGenerationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateMediaStorageConfigurationInput(v *UpdateMediaStorageConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMediaStorageConfigurationInput"}
	if v.ChannelARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelARN"))
	}
	if v.MediaStorageConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MediaStorageConfiguration"))
	} else if v.MediaStorageConfiguration != nil {
		if err := validateMediaStorageConfiguration(v.MediaStorageConfiguration); err != nil {
			invalidParams.AddNested("MediaStorageConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateNotificationConfigurationInput(v *UpdateNotificationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateNotificationConfigurationInput"}
	if v.NotificationConfiguration != nil {
		if err := validateNotificationConfiguration(v.NotificationConfiguration); err != nil {
			invalidParams.AddNested("NotificationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSignalingChannelInput(v *UpdateSignalingChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSignalingChannelInput"}
	if v.ChannelARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelARN"))
	}
	if v.CurrentVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CurrentVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateStreamInput(v *UpdateStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateStreamInput"}
	if v.CurrentVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CurrentVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
