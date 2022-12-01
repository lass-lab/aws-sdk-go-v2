// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains details about an activity that failed during an execution.
type ActivityFailedEventDetails struct {

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about an activity.
type ActivityListItem struct {

	// The Amazon Resource Name (ARN) that identifies the activity.
	//
	// This member is required.
	ActivityArn *string

	// The date the activity is created.
	//
	// This member is required.
	CreationDate *time.Time

	// The name of the activity. A name must not contain:
	//
	// * white space
	//
	// * brackets <
	// > { } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters " # % \ ^ | ~ ` $ & ,
	// ; : /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable logging with
	// CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and _.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Contains details about an activity scheduled during an execution.
type ActivityScheduledEventDetails struct {

	// The Amazon Resource Name (ARN) of the scheduled activity.
	//
	// This member is required.
	Resource *string

	// The maximum allowed duration between two heartbeats for the activity task.
	HeartbeatInSeconds *int64

	// The JSON data input to the activity task. Length constraints apply to the
	// payload size, and are expressed as bytes in UTF-8 encoding.
	Input *string

	// Contains details about the input for an execution history event.
	InputDetails *HistoryEventExecutionDataDetails

	// The maximum allowed duration of the activity task.
	TimeoutInSeconds *int64

	noSmithyDocumentSerde
}

// Contains details about an activity schedule failure that occurred during an
// execution.
type ActivityScheduleFailedEventDetails struct {

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about the start of an activity during an execution.
type ActivityStartedEventDetails struct {

	// The name of the worker that the task is assigned to. These names are provided by
	// the workers when calling GetActivityTask.
	WorkerName *string

	noSmithyDocumentSerde
}

// Contains details about an activity that successfully terminated during an
// execution.
type ActivitySucceededEventDetails struct {

	// The JSON data output by the activity task. Length constraints apply to the
	// payload size, and are expressed as bytes in UTF-8 encoding.
	Output *string

	// Contains details about the output of an execution history event.
	OutputDetails *HistoryEventExecutionDataDetails

	noSmithyDocumentSerde
}

// Contains details about an activity timeout that occurred during an execution.
type ActivityTimedOutEventDetails struct {

	// A more detailed explanation of the cause of the timeout.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// An object that describes workflow billing details.
type BillingDetails struct {

	// Billed duration of your workflow, in milliseconds.
	BilledDurationInMilliseconds int64

	// Billed memory consumption of your workflow, in MB.
	BilledMemoryUsedInMB int64

	noSmithyDocumentSerde
}

// Provides details about execution input or output.
type CloudWatchEventsExecutionDataDetails struct {

	// Indicates whether input or output was included in the response. Always true for
	// API calls.
	Included bool

	noSmithyDocumentSerde
}

type CloudWatchLogsLogGroup struct {

	// The ARN of the the CloudWatch log group to which you want your logs emitted to.
	// The ARN must end with :*
	LogGroupArn *string

	noSmithyDocumentSerde
}

// Contains details about an abort of an execution.
type ExecutionAbortedEventDetails struct {

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about an execution failure event.
type ExecutionFailedEventDetails struct {

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about an execution.
type ExecutionListItem struct {

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// This member is required.
	ExecutionArn *string

	// The name of the execution. A name must not contain:
	//
	// * white space
	//
	// * brackets <
	// > { } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters " # % \ ^ | ~ ` $ & ,
	// ; : /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable logging with
	// CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and _.
	//
	// This member is required.
	Name *string

	// The date the execution started.
	//
	// This member is required.
	StartDate *time.Time

	// The Amazon Resource Name (ARN) of the executed state machine.
	//
	// This member is required.
	StateMachineArn *string

	// The current status of the execution.
	//
	// This member is required.
	Status ExecutionStatus

	// The total number of items processed in a child workflow execution. This field is
	// returned only if mapRunArn was specified in the ListExecutions API action. If
	// stateMachineArn was specified in ListExecutions, the itemCount field isn't
	// returned.
	ItemCount *int32

	// The Amazon Resource Name (ARN) of a Map Run. This field is returned only if
	// mapRunArn was specified in the ListExecutions API action. If stateMachineArn was
	// specified in ListExecutions, the mapRunArn isn't returned.
	MapRunArn *string

	// If the execution already ended, the date the execution stopped.
	StopDate *time.Time

	noSmithyDocumentSerde
}

// Contains details about the start of the execution.
type ExecutionStartedEventDetails struct {

	// The JSON data input to the execution. Length constraints apply to the payload
	// size, and are expressed as bytes in UTF-8 encoding.
	Input *string

	// Contains details about the input for an execution history event.
	InputDetails *HistoryEventExecutionDataDetails

	// The Amazon Resource Name (ARN) of the IAM role used for executing Lambda tasks.
	RoleArn *string

	noSmithyDocumentSerde
}

// Contains details about the successful termination of the execution.
type ExecutionSucceededEventDetails struct {

	// The JSON data output by the execution. Length constraints apply to the payload
	// size, and are expressed as bytes in UTF-8 encoding.
	Output *string

	// Contains details about the output of an execution history event.
	OutputDetails *HistoryEventExecutionDataDetails

	noSmithyDocumentSerde
}

// Contains details about the execution timeout that occurred during the execution.
type ExecutionTimedOutEventDetails struct {

	// A more detailed explanation of the cause of the timeout.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about the events of an execution.
type HistoryEvent struct {

	// The id of the event. Events are numbered sequentially, starting at one.
	//
	// This member is required.
	Id int64

	// The date and time the event occurred.
	//
	// This member is required.
	Timestamp *time.Time

	// The type of the event.
	//
	// This member is required.
	Type HistoryEventType

	// Contains details about an activity that failed during an execution.
	ActivityFailedEventDetails *ActivityFailedEventDetails

	// Contains details about an activity schedule event that failed during an
	// execution.
	ActivityScheduleFailedEventDetails *ActivityScheduleFailedEventDetails

	// Contains details about an activity scheduled during an execution.
	ActivityScheduledEventDetails *ActivityScheduledEventDetails

	// Contains details about the start of an activity during an execution.
	ActivityStartedEventDetails *ActivityStartedEventDetails

	// Contains details about an activity that successfully terminated during an
	// execution.
	ActivitySucceededEventDetails *ActivitySucceededEventDetails

	// Contains details about an activity timeout that occurred during an execution.
	ActivityTimedOutEventDetails *ActivityTimedOutEventDetails

	// Contains details about an abort of an execution.
	ExecutionAbortedEventDetails *ExecutionAbortedEventDetails

	// Contains details about an execution failure event.
	ExecutionFailedEventDetails *ExecutionFailedEventDetails

	// Contains details about the start of the execution.
	ExecutionStartedEventDetails *ExecutionStartedEventDetails

	// Contains details about the successful termination of the execution.
	ExecutionSucceededEventDetails *ExecutionSucceededEventDetails

	// Contains details about the execution timeout that occurred during the execution.
	ExecutionTimedOutEventDetails *ExecutionTimedOutEventDetails

	// Contains details about a Lambda function that failed during an execution.
	LambdaFunctionFailedEventDetails *LambdaFunctionFailedEventDetails

	// Contains details about a failed Lambda function schedule event that occurred
	// during an execution.
	LambdaFunctionScheduleFailedEventDetails *LambdaFunctionScheduleFailedEventDetails

	// Contains details about a Lambda function scheduled during an execution.
	LambdaFunctionScheduledEventDetails *LambdaFunctionScheduledEventDetails

	// Contains details about a lambda function that failed to start during an
	// execution.
	LambdaFunctionStartFailedEventDetails *LambdaFunctionStartFailedEventDetails

	// Contains details about a Lambda function that terminated successfully during an
	// execution.
	LambdaFunctionSucceededEventDetails *LambdaFunctionSucceededEventDetails

	// Contains details about a Lambda function timeout that occurred during an
	// execution.
	LambdaFunctionTimedOutEventDetails *LambdaFunctionTimedOutEventDetails

	// Contains details about an iteration of a Map state that was aborted.
	MapIterationAbortedEventDetails *MapIterationEventDetails

	// Contains details about an iteration of a Map state that failed.
	MapIterationFailedEventDetails *MapIterationEventDetails

	// Contains details about an iteration of a Map state that was started.
	MapIterationStartedEventDetails *MapIterationEventDetails

	// Contains details about an iteration of a Map state that succeeded.
	MapIterationSucceededEventDetails *MapIterationEventDetails

	// Contains error and cause details about a Map Run that failed.
	MapRunFailedEventDetails *MapRunFailedEventDetails

	// Contains details, such as mapRunArn, and the start date and time of a Map Run.
	// mapRunArn is the Amazon Resource Name (ARN) of the Map Run that was started.
	MapRunStartedEventDetails *MapRunStartedEventDetails

	// Contains details about Map state that was started.
	MapStateStartedEventDetails *MapStateStartedEventDetails

	// The id of the previous event.
	PreviousEventId int64

	// Contains details about a state entered during an execution.
	StateEnteredEventDetails *StateEnteredEventDetails

	// Contains details about an exit from a state during an execution.
	StateExitedEventDetails *StateExitedEventDetails

	// Contains details about the failure of a task.
	TaskFailedEventDetails *TaskFailedEventDetails

	// Contains details about a task that was scheduled.
	TaskScheduledEventDetails *TaskScheduledEventDetails

	// Contains details about a task that failed to start.
	TaskStartFailedEventDetails *TaskStartFailedEventDetails

	// Contains details about a task that was started.
	TaskStartedEventDetails *TaskStartedEventDetails

	// Contains details about a task that where the submit failed.
	TaskSubmitFailedEventDetails *TaskSubmitFailedEventDetails

	// Contains details about a submitted task.
	TaskSubmittedEventDetails *TaskSubmittedEventDetails

	// Contains details about a task that succeeded.
	TaskSucceededEventDetails *TaskSucceededEventDetails

	// Contains details about a task that timed out.
	TaskTimedOutEventDetails *TaskTimedOutEventDetails

	noSmithyDocumentSerde
}

// Provides details about input or output in an execution history event.
type HistoryEventExecutionDataDetails struct {

	// Indicates whether input or output was truncated in the response. Always false
	// for API calls.
	Truncated bool

	noSmithyDocumentSerde
}

// Contains details about a Lambda function that failed during an execution.
type LambdaFunctionFailedEventDetails struct {

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about a Lambda function scheduled during an execution.
type LambdaFunctionScheduledEventDetails struct {

	// The Amazon Resource Name (ARN) of the scheduled Lambda function.
	//
	// This member is required.
	Resource *string

	// The JSON data input to the Lambda function. Length constraints apply to the
	// payload size, and are expressed as bytes in UTF-8 encoding.
	Input *string

	// Contains details about input for an execution history event.
	InputDetails *HistoryEventExecutionDataDetails

	// The credentials that Step Functions uses for the task.
	TaskCredentials *TaskCredentials

	// The maximum allowed duration of the Lambda function.
	TimeoutInSeconds *int64

	noSmithyDocumentSerde
}

// Contains details about a failed Lambda function schedule event that occurred
// during an execution.
type LambdaFunctionScheduleFailedEventDetails struct {

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about a lambda function that failed to start during an
// execution.
type LambdaFunctionStartFailedEventDetails struct {

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about a Lambda function that successfully terminated during an
// execution.
type LambdaFunctionSucceededEventDetails struct {

	// The JSON data output by the Lambda function. Length constraints apply to the
	// payload size, and are expressed as bytes in UTF-8 encoding.
	Output *string

	// Contains details about the output of an execution history event.
	OutputDetails *HistoryEventExecutionDataDetails

	noSmithyDocumentSerde
}

// Contains details about a Lambda function timeout that occurred during an
// execution.
type LambdaFunctionTimedOutEventDetails struct {

	// A more detailed explanation of the cause of the timeout.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

type LogDestination struct {

	// An object describing a CloudWatch log group. For more information, see
	// AWS::Logs::LogGroup
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html)
	// in the CloudFormation User Guide.
	CloudWatchLogsLogGroup *CloudWatchLogsLogGroup

	noSmithyDocumentSerde
}

// The LoggingConfiguration data type is used to set CloudWatch Logs options.
type LoggingConfiguration struct {

	// An array of objects that describes where your execution history events will be
	// logged. Limited to size 1. Required, if your log level is not set to OFF.
	Destinations []LogDestination

	// Determines whether execution data is included in your log. When set to false,
	// data is excluded.
	IncludeExecutionData bool

	// Defines which category of execution history events are logged.
	Level LogLevel

	noSmithyDocumentSerde
}

// Contains details about an iteration of a Map state.
type MapIterationEventDetails struct {

	// The index of the array belonging to the Map state iteration.
	Index int32

	// The name of the iteration’s parent Map state.
	Name *string

	noSmithyDocumentSerde
}

// Contains details about all of the child workflow executions started by a Map
// Run.
type MapRunExecutionCounts struct {

	// The total number of child workflow executions that were started by a Map Run and
	// were running, but were either stopped by the user or by Step Functions because
	// the Map Run failed.
	//
	// This member is required.
	Aborted int64

	// The total number of child workflow executions that were started by a Map Run,
	// but have failed.
	//
	// This member is required.
	Failed int64

	// The total number of child workflow executions that were started by a Map Run,
	// but haven't started executing yet.
	//
	// This member is required.
	Pending int64

	// Returns the count of child workflow executions whose results were written by
	// ResultWriter. For more information, see ResultWriter
	// (https://docs.aws.amazon.com/step-functions/latest/dg/input-output-resultwriter.html)
	// in the Step Functions Developer Guide.
	//
	// This member is required.
	ResultsWritten int64

	// The total number of child workflow executions that were started by a Map Run and
	// are currently in-progress.
	//
	// This member is required.
	Running int64

	// The total number of child workflow executions that were started by a Map Run and
	// have completed successfully.
	//
	// This member is required.
	Succeeded int64

	// The total number of child workflow executions that were started by a Map Run and
	// have timed out.
	//
	// This member is required.
	TimedOut int64

	// The total number of child workflow executions that were started by a Map Run.
	//
	// This member is required.
	Total int64

	noSmithyDocumentSerde
}

// Contains details about a Map Run failure event that occurred during a state
// machine execution.
type MapRunFailedEventDetails struct {

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the Map Run failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about items that were processed in all of the child workflow
// executions that were started by a Map Run.
type MapRunItemCounts struct {

	// The total number of items processed in child workflow executions that were
	// either stopped by the user or by Step Functions, because the Map Run failed.
	//
	// This member is required.
	Aborted int64

	// The total number of items processed in child workflow executions that have
	// failed.
	//
	// This member is required.
	Failed int64

	// The total number of items to process in child workflow executions that haven't
	// started running yet.
	//
	// This member is required.
	Pending int64

	// Returns the count of items whose results were written by ResultWriter. For more
	// information, see ResultWriter
	// (https://docs.aws.amazon.com/step-functions/latest/dg/input-output-resultwriter.html)
	// in the Step Functions Developer Guide.
	//
	// This member is required.
	ResultsWritten int64

	// The total number of items being processed in child workflow executions that are
	// currently in-progress.
	//
	// This member is required.
	Running int64

	// The total number of items processed in child workflow executions that have
	// completed successfully.
	//
	// This member is required.
	Succeeded int64

	// The total number of items processed in child workflow executions that have timed
	// out.
	//
	// This member is required.
	TimedOut int64

	// The total number of items processed in all the child workflow executions started
	// by a Map Run.
	//
	// This member is required.
	Total int64

	noSmithyDocumentSerde
}

// Contains details about a specific Map Run.
type MapRunListItem struct {

	// The executionArn of the execution from which the Map Run was started.
	//
	// This member is required.
	ExecutionArn *string

	// The Amazon Resource Name (ARN) of the Map Run.
	//
	// This member is required.
	MapRunArn *string

	// The date on which the Map Run started.
	//
	// This member is required.
	StartDate *time.Time

	// The Amazon Resource Name (ARN) of the executed state machine.
	//
	// This member is required.
	StateMachineArn *string

	// The date on which the Map Run stopped.
	StopDate *time.Time

	noSmithyDocumentSerde
}

// Contains details about a Map Run that was started during a state machine
// execution.
type MapRunStartedEventDetails struct {

	// The Amazon Resource Name (ARN) of a Map Run that was started.
	MapRunArn *string

	noSmithyDocumentSerde
}

// Details about a Map state that was started.
type MapStateStartedEventDetails struct {

	// The size of the array for Map state iterations.
	Length int32

	noSmithyDocumentSerde
}

// Contains details about a state entered during an execution.
type StateEnteredEventDetails struct {

	// The name of the state.
	//
	// This member is required.
	Name *string

	// The string that contains the JSON input data for the state. Length constraints
	// apply to the payload size, and are expressed as bytes in UTF-8 encoding.
	Input *string

	// Contains details about the input for an execution history event.
	InputDetails *HistoryEventExecutionDataDetails

	noSmithyDocumentSerde
}

// Contains details about an exit from a state during an execution.
type StateExitedEventDetails struct {

	// The name of the state. A name must not contain:
	//
	// * white space
	//
	// * brackets < > {
	// } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters " # % \ ^ | ~ ` $ & , ; :
	// /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable logging with
	// CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and _.
	//
	// This member is required.
	Name *string

	// The JSON output data of the state. Length constraints apply to the payload size,
	// and are expressed as bytes in UTF-8 encoding.
	Output *string

	// Contains details about the output of an execution history event.
	OutputDetails *HistoryEventExecutionDataDetails

	noSmithyDocumentSerde
}

// Contains details about the state machine.
type StateMachineListItem struct {

	// The date the state machine is created.
	//
	// This member is required.
	CreationDate *time.Time

	// The name of the state machine. A name must not contain:
	//
	// * white space
	//
	// *
	// brackets < > { } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters " # % \ ^
	// | ~ ` $ & , ; : /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable
	// logging with CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and
	// _.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) that identifies the state machine.
	//
	// This member is required.
	StateMachineArn *string

	//
	//
	// This member is required.
	Type StateMachineType

	noSmithyDocumentSerde
}

// Tags are key-value pairs that can be associated with Step Functions state
// machines and activities. An array of key-value pairs. For more information, see
// Using Cost Allocation Tags
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the Amazon Web Services Billing and Cost Management User Guide, and
// Controlling Access Using IAM Tags
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html). Tags
// may only contain Unicode letters, digits, white space, or these symbols: _ . : /
// = + - @.
type Tag struct {

	// The key of a tag.
	Key *string

	// The value of a tag.
	Value *string

	noSmithyDocumentSerde
}

// Contains details about the credentials that Step Functions uses for a task.
type TaskCredentials struct {

	// The ARN of an IAM role that Step Functions assumes for the task. The role can
	// allow cross-account access to resources.
	RoleArn *string

	noSmithyDocumentSerde
}

// Contains details about a task failure event.
type TaskFailedEventDetails struct {

	// The action of the resource called by a task state.
	//
	// This member is required.
	Resource *string

	// The service name of the resource in a task state.
	//
	// This member is required.
	ResourceType *string

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about a task scheduled during an execution.
type TaskScheduledEventDetails struct {

	// The JSON data passed to the resource referenced in a task state. Length
	// constraints apply to the payload size, and are expressed as bytes in UTF-8
	// encoding.
	//
	// This member is required.
	Parameters *string

	// The region of the scheduled task
	//
	// This member is required.
	Region *string

	// The action of the resource called by a task state.
	//
	// This member is required.
	Resource *string

	// The service name of the resource in a task state.
	//
	// This member is required.
	ResourceType *string

	// The maximum allowed duration between two heartbeats for the task.
	HeartbeatInSeconds *int64

	// The credentials that Step Functions uses for the task.
	TaskCredentials *TaskCredentials

	// The maximum allowed duration of the task.
	TimeoutInSeconds *int64

	noSmithyDocumentSerde
}

// Contains details about the start of a task during an execution.
type TaskStartedEventDetails struct {

	// The action of the resource called by a task state.
	//
	// This member is required.
	Resource *string

	// The service name of the resource in a task state.
	//
	// This member is required.
	ResourceType *string

	noSmithyDocumentSerde
}

// Contains details about a task that failed to start during an execution.
type TaskStartFailedEventDetails struct {

	// The action of the resource called by a task state.
	//
	// This member is required.
	Resource *string

	// The service name of the resource in a task state.
	//
	// This member is required.
	ResourceType *string

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about a task that failed to submit during an execution.
type TaskSubmitFailedEventDetails struct {

	// The action of the resource called by a task state.
	//
	// This member is required.
	Resource *string

	// The service name of the resource in a task state.
	//
	// This member is required.
	ResourceType *string

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Contains details about a task submitted to a resource .
type TaskSubmittedEventDetails struct {

	// The action of the resource called by a task state.
	//
	// This member is required.
	Resource *string

	// The service name of the resource in a task state.
	//
	// This member is required.
	ResourceType *string

	// The response from a resource when a task has started. Length constraints apply
	// to the payload size, and are expressed as bytes in UTF-8 encoding.
	Output *string

	// Contains details about the output of an execution history event.
	OutputDetails *HistoryEventExecutionDataDetails

	noSmithyDocumentSerde
}

// Contains details about the successful completion of a task state.
type TaskSucceededEventDetails struct {

	// The action of the resource called by a task state.
	//
	// This member is required.
	Resource *string

	// The service name of the resource in a task state.
	//
	// This member is required.
	ResourceType *string

	// The full JSON response from a resource when a task has succeeded. This response
	// becomes the output of the related task. Length constraints apply to the payload
	// size, and are expressed as bytes in UTF-8 encoding.
	Output *string

	// Contains details about the output of an execution history event.
	OutputDetails *HistoryEventExecutionDataDetails

	noSmithyDocumentSerde
}

// Contains details about a resource timeout that occurred during an execution.
type TaskTimedOutEventDetails struct {

	// The action of the resource called by a task state.
	//
	// This member is required.
	Resource *string

	// The service name of the resource in a task state.
	//
	// This member is required.
	ResourceType *string

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	noSmithyDocumentSerde
}

// Selects whether or not the state machine's X-Ray tracing is enabled. Default is
// false
type TracingConfiguration struct {

	// When set to true, X-Ray tracing is enabled.
	Enabled bool

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
