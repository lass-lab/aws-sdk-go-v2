// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores a DB cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target DB cluster is created from the source DB cluster with the same
// configuration as the original DB cluster, except that the new DB cluster is
// created with the default DB security group. This action only restores the DB
// cluster, not the DB instances for that DB cluster. You must invoke the
// CreateDBInstance action to create DB instances for the restored DB cluster,
// specifying the identifier of the restored DB cluster in DBClusterIdentifier. You
// can create DB instances only after the RestoreDBClusterToPointInTime action has
// completed and the DB cluster is available. For more information on Amazon
// Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) RestoreDBClusterToPointInTime(ctx context.Context, params *RestoreDBClusterToPointInTimeInput, optFns ...func(*Options)) (*RestoreDBClusterToPointInTimeOutput, error) {
	if params == nil {
		params = &RestoreDBClusterToPointInTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBClusterToPointInTime", params, optFns, addOperationRestoreDBClusterToPointInTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBClusterToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RestoreDBClusterToPointInTimeInput struct {

	// The name of the new DB cluster to be created. Constraints:
	//
	// * Must contain from
	// 1 to 63 letters, numbers, or hyphens
	//
	// * First character must be a letter
	//
	// *
	// Can't end with a hyphen or contain two consecutive hyphens
	//
	// This member is required.
	DBClusterIdentifier *string

	// The identifier of the source DB cluster from which to restore. Constraints:
	//
	// *
	// Must match the identifier of an existing DBCluster.
	//
	// This member is required.
	SourceDBClusterIdentifier *string

	// The target backtrack window, in seconds. To disable backtracking, set this value
	// to 0. Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	// Default: 0 Constraints:
	//
	// * If specified, this value must be set to a number from
	// 0 to 259,200 (72 hours).
	BacktrackWindow *int64

	// A value that indicates whether to copy all tags from the restored DB cluster to
	// snapshots of the restored DB cluster. The default is not to copy them.
	CopyTagsToSnapshot *bool

	// The name of the DB cluster parameter group to associate with this DB cluster. If
	// this argument is omitted, the default DB cluster parameter group for the
	// specified engine is used. Constraints:
	//
	// * If supplied, must match the name of an
	// existing DB cluster parameter group.
	//
	// * Must be 1 to 255 letters, numbers, or
	// hyphens.
	//
	// * First character must be a letter.
	//
	// * Can't end with a hyphen or
	// contain two consecutive hyphens.
	DBClusterParameterGroupName *string

	// The DB subnet group name to use for the new DB cluster. Constraints: If
	// supplied, must match the name of an existing DBSubnetGroup. Example:
	// mySubnetgroup
	DBSubnetGroupName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool

	// Specify the Active Directory directory ID to restore the DB cluster in. The
	// domain must be created prior to this operation. For Amazon Aurora DB clusters,
	// Amazon RDS can use Kerberos Authentication to authenticate users that connect to
	// the DB cluster. For more information, see Kerberos Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html)
	// in the Amazon Aurora User Guide.
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string

	// The list of logs that the restored DB cluster is to export to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. For more information,
	// see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	EnableCloudwatchLogsExports []string

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// For more information, see  IAM Database Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide.
	EnableIAMDatabaseAuthentication *bool

	// The AWS KMS key identifier to use when restoring an encrypted DB cluster from an
	// encrypted DB cluster. The AWS KMS key identifier is the key ARN, key ID, alias
	// ARN, or alias name for the AWS KMS customer master key (CMK). To use a CMK in a
	// different AWS account, specify the key ARN or alias ARN. You can restore to a
	// new DB cluster and encrypt the new DB cluster with a AWS KMS CMK that is
	// different than the AWS KMS key used to encrypt the source DB cluster. The new DB
	// cluster is encrypted with the AWS KMS CMK identified by the KmsKeyId parameter.
	// If you don't specify a value for the KmsKeyId parameter, then the following
	// occurs:
	//
	// * If the DB cluster is encrypted, then the restored DB cluster is
	// encrypted using the AWS KMS CMK that was used to encrypt the source DB
	// cluster.
	//
	// * If the DB cluster isn't encrypted, then the restored DB cluster
	// isn't encrypted.
	//
	// If DBClusterIdentifier refers to a DB cluster that isn't
	// encrypted, then the restore request is rejected.
	KmsKeyId *string

	// The name of the option group for the new DB cluster.
	OptionGroupName *string

	// The port number on which the new DB cluster accepts connections. Constraints: A
	// value from 1150-65535. Default: The default port for the engine.
	Port *int32

	// The date and time to restore the DB cluster to. Valid Values: Value must be a
	// time in Universal Coordinated Time (UTC) format Constraints:
	//
	// * Must be before
	// the latest restorable time for the DB instance
	//
	// * Must be specified if
	// UseLatestRestorableTime parameter isn't provided
	//
	// * Can't be specified if the
	// UseLatestRestorableTime parameter is enabled
	//
	// * Can't be specified if the
	// RestoreType parameter is copy-on-write
	//
	// Example: 2015-03-07T23:45:00Z
	RestoreToTime *time.Time

	// The type of restore to be performed. You can specify one of the following
	// values:
	//
	// * full-copy - The new DB cluster is restored as a full copy of the
	// source DB cluster.
	//
	// * copy-on-write - The new DB cluster is restored as a clone
	// of the source DB cluster.
	//
	// Constraints: You can't specify copy-on-write if the
	// engine version of the source DB cluster is earlier than 1.11. If you don't
	// specify a RestoreType value, then the new DB cluster is restored as a full copy
	// of the source DB cluster.
	RestoreType *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []types.Tag

	// A value that indicates whether to restore the DB cluster to the latest
	// restorable backup time. By default, the DB cluster isn't restored to the latest
	// restorable backup time. Constraints: Can't be specified if RestoreToTime
	// parameter is provided.
	UseLatestRestorableTime bool

	// A list of VPC security groups that the new DB cluster belongs to.
	VpcSecurityGroupIds []string
}

type RestoreDBClusterToPointInTimeOutput struct {

	// Contains the details of an Amazon Aurora DB cluster. This data type is used as a
	// response element in the DescribeDBClusters, StopDBCluster, and StartDBCluster
	// actions.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRestoreDBClusterToPointInTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterToPointInTime{}, middleware.After)
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
	if err = addOpRestoreDBClusterToPointInTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterToPointInTime(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBClusterToPointInTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBClusterToPointInTime",
	}
}
