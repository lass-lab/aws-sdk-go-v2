// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Details of an EC2 AMI.
type Ami struct {

	// The account ID of the owner of the AMI.
	AccountId *string

	// The description of the EC2 AMI. Minimum and maximum length are in characters.
	Description *string

	// The AMI ID of the EC2 AMI.
	Image *string

	// The name of the EC2 AMI.
	Name *string

	// The AWS Region of the EC2 AMI.
	Region *string

	// Image state shows the image status and the reason for that status.
	State *ImageState
}

// Define and configure the output AMIs of the pipeline.
type AmiDistributionConfiguration struct {

	// The tags to apply to AMIs distributed to this Region.
	AmiTags map[string]string

	// The description of the distribution configuration. Minimum and maximum length
	// are in characters.
	Description *string

	// The KMS key identifier used to encrypt the distributed image.
	KmsKeyId *string

	// Launch permissions can be used to configure which AWS accounts can use the AMI
	// to launch instances.
	LaunchPermission *LaunchPermissionConfiguration

	// The name of the distribution configuration.
	Name *string

	// The ID of an account to which you want to distribute an image.
	TargetAccountIds []string
}

// A detailed view of a component.
type Component struct {

	// The Amazon Resource Name (ARN) of the component.
	Arn *string

	// The change description of the component.
	ChangeDescription *string

	// The data of the component.
	Data *string

	// The date that the component was created.
	DateCreated *string

	// The description of the component.
	Description *string

	// The encryption status of the component.
	Encrypted bool

	// The KMS key identifier used to encrypt the component.
	KmsKeyId *string

	// The name of the component.
	Name *string

	// The owner of the component.
	Owner *string

	// The platform of the component.
	Platform Platform

	// The operating system (OS) version supported by the component. If the OS
	// information is available, a prefix match is performed against the parent image
	// OS version during image recipe creation.
	SupportedOsVersions []string

	// The tags associated with the component.
	Tags map[string]string

	// The type of the component denotes whether the component is used to build the
	// image or only to test it.
	Type ComponentType

	// The version of the component.
	Version *string
}

// Configuration details of the component.
type ComponentConfiguration struct {

	// The Amazon Resource Name (ARN) of the component.
	//
	// This member is required.
	ComponentArn *string
}

// A high-level summary of a component.
type ComponentSummary struct {

	// The Amazon Resource Name (ARN) of the component.
	Arn *string

	// The change description of the component.
	ChangeDescription *string

	// The date that the component was created.
	DateCreated *string

	// The description of the component.
	Description *string

	// The name of the component.
	Name *string

	// The owner of the component.
	Owner *string

	// The platform of the component.
	Platform Platform

	// The operating system (OS) version supported by the component. If the OS
	// information is available, a prefix match is performed against the parent image
	// OS version during image recipe creation.
	SupportedOsVersions []string

	// The tags associated with the component.
	Tags map[string]string

	// The type of the component denotes whether the component is used to build the
	// image or only to test it.
	Type ComponentType

	// The version of the component.
	Version *string
}

// A high-level overview of a component semantic version.
type ComponentVersion struct {

	// The Amazon Resource Name (ARN) of the component.
	Arn *string

	// The date that the component was created.
	DateCreated *string

	// The description of the component.
	Description *string

	// The name of the component.
	Name *string

	// The owner of the component.
	Owner *string

	// The platform of the component.
	Platform Platform

	// The operating system (OS) version supported by the component. If the OS
	// information is available, a prefix match is performed against the parent image
	// OS version during image recipe creation.
	SupportedOsVersions []string

	// The type of the component denotes whether the component is used to build the
	// image or only to test it.
	Type ComponentType

	// The semantic version of the component.
	Version *string
}

// Defines the settings for a specific Region.
type Distribution struct {

	// The target Region.
	//
	// This member is required.
	Region *string

	// The specific AMI settings (for example, launch permissions, AMI tags).
	AmiDistributionConfiguration *AmiDistributionConfiguration

	// The License Manager Configuration to associate with the AMI in the specified
	// Region.
	LicenseConfigurationArns []string
}

// A distribution configuration.
type DistributionConfiguration struct {

	// The maximum duration in minutes for this distribution configuration.
	//
	// This member is required.
	TimeoutMinutes int32

	// The Amazon Resource Name (ARN) of the distribution configuration.
	Arn *string

	// The date on which this distribution configuration was created.
	DateCreated *string

	// The date on which this distribution configuration was last updated.
	DateUpdated *string

	// The description of the distribution configuration.
	Description *string

	// The distributions of the distribution configuration.
	Distributions []Distribution

	// The name of the distribution configuration.
	Name *string

	// The tags of the distribution configuration.
	Tags map[string]string
}

// A high-level overview of a distribution configuration.
type DistributionConfigurationSummary struct {

	// The Amazon Resource Name (ARN) of the distribution configuration.
	Arn *string

	// The date on which the distribution configuration was created.
	DateCreated *string

	// The date on which the distribution configuration was updated.
	DateUpdated *string

	// The description of the distribution configuration.
	Description *string

	// The name of the distribution configuration.
	Name *string

	// The tags associated with the distribution configuration.
	Tags map[string]string
}

// Amazon EBS-specific block device mapping specifications.
type EbsInstanceBlockDeviceSpecification struct {

	// Use to configure delete on termination of the associated device.
	DeleteOnTermination bool

	// Use to configure device encryption.
	Encrypted bool

	// Use to configure device IOPS.
	Iops int32

	// Use to configure the KMS key to use when encrypting the device.
	KmsKeyId *string

	// The snapshot that defines the device contents.
	SnapshotId *string

	// Use to override the device's volume size.
	VolumeSize int32

	// Use to override the device's volume type.
	VolumeType EbsVolumeType
}

// A filter name and value pair that is used to return a more specific list of
// results from a list operation. Filters can be used to match a set of resources
// by specific criteria, such as tags, attributes, or IDs.
type Filter struct {

	// The name of the filter. Filter names are case-sensitive.
	Name *string

	// The filter values. Filter values are case-sensitive.
	Values []string
}

// An image build version.
type Image struct {

	// The Amazon Resource Name (ARN) of the image.
	Arn *string

	// The date on which this image was created.
	DateCreated *string

	// The distribution configuration used when creating this image.
	DistributionConfiguration *DistributionConfiguration

	// Collects additional information about the image being created, including the
	// operating system (OS) version and package list. This information is used to
	// enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled bool

	// The image recipe used when creating the image.
	ImageRecipe *ImageRecipe

	// The image tests configuration used when creating this image.
	ImageTestsConfiguration *ImageTestsConfiguration

	// The infrastructure used when creating this image.
	InfrastructureConfiguration *InfrastructureConfiguration

	// The name of the image.
	Name *string

	// The operating system version of the instance. For example, Amazon Linux 2,
	// Ubuntu 18, or Microsoft Windows Server 2019.
	OsVersion *string

	// The output resources produced when creating this image.
	OutputResources *OutputResources

	// The platform of the image.
	Platform Platform

	// The Amazon Resource Name (ARN) of the image pipeline that created this image.
	SourcePipelineArn *string

	// The name of the image pipeline that created this image.
	SourcePipelineName *string

	// The state of the image.
	State *ImageState

	// The tags of the image.
	Tags map[string]string

	// The semantic version of the image.
	Version *string
}

// Details of an image pipeline.
type ImagePipeline struct {

	// The Amazon Resource Name (ARN) of the image pipeline.
	Arn *string

	// The date on which this image pipeline was created.
	DateCreated *string

	// The date on which this image pipeline was last run.
	DateLastRun *string

	// The date on which this image pipeline will next be run.
	DateNextRun *string

	// The date on which this image pipeline was last updated.
	DateUpdated *string

	// The description of the image pipeline.
	Description *string

	// The Amazon Resource Name (ARN) of the distribution configuration associated with
	// this image pipeline.
	DistributionConfigurationArn *string

	// Collects additional information about the image being created, including the
	// operating system (OS) version and package list. This information is used to
	// enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled bool

	// The Amazon Resource Name (ARN) of the image recipe associated with this image
	// pipeline.
	ImageRecipeArn *string

	// The image tests configuration of the image pipeline.
	ImageTestsConfiguration *ImageTestsConfiguration

	// The Amazon Resource Name (ARN) of the infrastructure configuration associated
	// with this image pipeline.
	InfrastructureConfigurationArn *string

	// The name of the image pipeline.
	Name *string

	// The platform of the image pipeline.
	Platform Platform

	// The schedule of the image pipeline.
	Schedule *Schedule

	// The status of the image pipeline.
	Status PipelineStatus

	// The tags of this image pipeline.
	Tags map[string]string
}

// An image recipe.
type ImageRecipe struct {

	// The Amazon Resource Name (ARN) of the image recipe.
	Arn *string

	// The block device mappings to apply when creating images from this recipe.
	BlockDeviceMappings []InstanceBlockDeviceMapping

	// The components of the image recipe.
	Components []ComponentConfiguration

	// The date on which this image recipe was created.
	DateCreated *string

	// The description of the image recipe.
	Description *string

	// The name of the image recipe.
	Name *string

	// The owner of the image recipe.
	Owner *string

	// The parent image of the image recipe.
	ParentImage *string

	// The platform of the image recipe.
	Platform Platform

	// The tags of the image recipe.
	Tags map[string]string

	// The version of the image recipe.
	Version *string

	// The working directory to be used during build and test workflows.
	WorkingDirectory *string
}

// A summary of an image recipe.
type ImageRecipeSummary struct {

	// The Amazon Resource Name (ARN) of the image recipe.
	Arn *string

	// The date on which this image recipe was created.
	DateCreated *string

	// The name of the image recipe.
	Name *string

	// The owner of the image recipe.
	Owner *string

	// The parent image of the image recipe.
	ParentImage *string

	// The platform of the image recipe.
	Platform Platform

	// The tags of the image recipe.
	Tags map[string]string
}

// Image state shows the image status and the reason for that status.
type ImageState struct {

	// The reason for the image's status.
	Reason *string

	// The status of the image.
	Status ImageStatus
}

// An image summary.
type ImageSummary struct {

	// The Amazon Resource Name (ARN) of the image.
	Arn *string

	// The date on which this image was created.
	DateCreated *string

	// The name of the image.
	Name *string

	// The operating system version of the instance. For example, Amazon Linux 2,
	// Ubuntu 18, or Microsoft Windows Server 2019.
	OsVersion *string

	// The output resources produced when creating this image.
	OutputResources *OutputResources

	// The owner of the image.
	Owner *string

	// The platform of the image.
	Platform Platform

	// The state of the image.
	State *ImageState

	// The tags of the image.
	Tags map[string]string

	// The version of the image.
	Version *string
}

// Image tests configuration.
type ImageTestsConfiguration struct {

	// Defines if tests should be executed when building this image.
	ImageTestsEnabled bool

	// The maximum time in minutes that tests are permitted to run.
	TimeoutMinutes int32
}

// An image semantic version.
type ImageVersion struct {

	// The Amazon Resource Name (ARN) of the image semantic version.
	Arn *string

	// The date at which this image semantic version was created.
	DateCreated *string

	// The name of the image semantic version.
	Name *string

	// The operating system version of the instance. For example, Amazon Linux 2,
	// Ubuntu 18, or Microsoft Windows Server 2019.
	OsVersion *string

	// The owner of the image semantic version.
	Owner *string

	// The platform of the image semantic version.
	Platform Platform

	// The semantic version of the image semantic version.
	Version *string
}

// Details of the infrastructure configuration.
type InfrastructureConfiguration struct {

	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn *string

	// The date on which the infrastructure configuration was created.
	DateCreated *string

	// The date on which the infrastructure configuration was last updated.
	DateUpdated *string

	// The description of the infrastructure configuration.
	Description *string

	// The instance profile of the infrastructure configuration.
	InstanceProfileName *string

	// The instance types of the infrastructure configuration.
	InstanceTypes []string

	// The EC2 key pair of the infrastructure configuration.
	KeyPair *string

	// The logging configuration of the infrastructure configuration.
	Logging *Logging

	// The name of the infrastructure configuration.
	Name *string

	// The tags attached to the resource created by Image Builder.
	ResourceTags map[string]string

	// The security group IDs of the infrastructure configuration.
	SecurityGroupIds []string

	// The SNS topic Amazon Resource Name (ARN) of the infrastructure configuration.
	SnsTopicArn *string

	// The subnet ID of the infrastructure configuration.
	SubnetId *string

	// The tags of the infrastructure configuration.
	Tags map[string]string

	// The terminate instance on failure configuration of the infrastructure
	// configuration.
	TerminateInstanceOnFailure bool
}

// The infrastructure used when building EC2 AMIs.
type InfrastructureConfigurationSummary struct {

	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn *string

	// The date on which the infrastructure configuration was created.
	DateCreated *string

	// The date on which the infrastructure configuration was last updated.
	DateUpdated *string

	// The description of the infrastructure configuration.
	Description *string

	// The name of the infrastructure configuration.
	Name *string

	// The tags attached to the image created by Image Builder.
	ResourceTags map[string]string

	// The tags of the infrastructure configuration.
	Tags map[string]string
}

// Defines block device mappings for the instance used to configure your image.
type InstanceBlockDeviceMapping struct {

	// The device to which these mappings apply.
	DeviceName *string

	// Use to manage Amazon EBS-specific configuration for this mapping.
	Ebs *EbsInstanceBlockDeviceSpecification

	// Use to remove a mapping from the parent image.
	NoDevice *string

	// Use to manage instance ephemeral devices.
	VirtualName *string
}

// Describes the configuration for a launch permission. The launch permission
// modification request is sent to the EC2 ModifyImageAttribute
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyImageAttribute.html)
// API on behalf of the user for each Region they have selected to distribute the
// AMI. To make an AMI public, set the launch permission authorized accounts to
// all. See the examples for making an AMI public at EC2 ModifyImageAttribute
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyImageAttribute.html).
type LaunchPermissionConfiguration struct {

	// The name of the group.
	UserGroups []string

	// The AWS account ID.
	UserIds []string
}

// Logging configuration defines where Image Builder uploads your logs.
type Logging struct {

	// The Amazon S3 logging configuration.
	S3Logs *S3Logs
}

// The resources produced by this image.
type OutputResources struct {

	// The EC2 AMIs created by this image.
	Amis []Ami
}

// Amazon S3 logging configuration.
type S3Logs struct {

	// The Amazon S3 bucket in which to store the logs.
	S3BucketName *string

	// The Amazon S3 path in which to store the logs.
	S3KeyPrefix *string
}

// A schedule configures how often and when a pipeline will automatically create a
// new image.
type Schedule struct {

	// The condition configures when the pipeline should trigger a new image build.
	// When the pipelineExecutionStartCondition is set to
	// EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE, and you use semantic version
	// filters on the source image or components in your image recipe, EC2 Image
	// Builder will build a new image only when there are new versions of the image or
	// components in your recipe that match the semantic version filter. When it is set
	// to EXPRESSION_MATCH_ONLY, it will build a new image every time the CRON
	// expression matches the current time. For semantic version syntax, see
	// CreateComponent
	// (https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_CreateComponent.html)
	// in the EC2 Image Builder API Reference.
	PipelineExecutionStartCondition PipelineExecutionStartCondition

	// The cron expression determines how often EC2 Image Builder evaluates your
	// pipelineExecutionStartCondition. For information on how to format a cron
	// expression in Image Builder, see Use cron expressions in EC2 Image Builder
	// (https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-cron.html).
	ScheduleExpression *string
}
