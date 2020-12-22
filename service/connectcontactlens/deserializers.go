// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcontactlens

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/connectcontactlens/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
)

type awsRestjson1_deserializeOpListRealtimeContactAnalysisSegments struct {
}

func (*awsRestjson1_deserializeOpListRealtimeContactAnalysisSegments) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpListRealtimeContactAnalysisSegments) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorListRealtimeContactAnalysisSegments(response, &metadata)
	}
	output := &ListRealtimeContactAnalysisSegmentsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentListRealtimeContactAnalysisSegmentsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorListRealtimeContactAnalysisSegments(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("AccessDeniedException", errorCode):
		return awsRestjson1_deserializeErrorAccessDeniedException(response, errorBody)

	case strings.EqualFold("InternalServiceException", errorCode):
		return awsRestjson1_deserializeErrorInternalServiceException(response, errorBody)

	case strings.EqualFold("InvalidRequestException", errorCode):
		return awsRestjson1_deserializeErrorInvalidRequestException(response, errorBody)

	case strings.EqualFold("ResourceNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorResourceNotFoundException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentListRealtimeContactAnalysisSegmentsOutput(v **ListRealtimeContactAnalysisSegmentsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *ListRealtimeContactAnalysisSegmentsOutput
	if *v == nil {
		sv = &ListRealtimeContactAnalysisSegmentsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "NextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected NextToken to be of type string, got %T instead", value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		case "Segments":
			if err := awsRestjson1_deserializeDocumentRealtimeContactAnalysisSegments(&sv.Segments, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorAccessDeniedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.AccessDeniedException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentAccessDeniedException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInternalServiceException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalServiceException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInternalServiceException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInvalidRequestException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InvalidRequestException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInvalidRequestException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorResourceNotFoundException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ResourceNotFoundException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentResourceNotFoundException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorThrottlingException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ThrottlingException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentThrottlingException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocumentAccessDeniedException(v **types.AccessDeniedException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.AccessDeniedException
	if *v == nil {
		sv = &types.AccessDeniedException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentCategories(v **types.Categories, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Categories
	if *v == nil {
		sv = &types.Categories{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "MatchedCategories":
			if err := awsRestjson1_deserializeDocumentMatchedCategories(&sv.MatchedCategories, value); err != nil {
				return err
			}

		case "MatchedDetails":
			if err := awsRestjson1_deserializeDocumentMatchedDetails(&sv.MatchedDetails, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentCategoryDetails(v **types.CategoryDetails, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.CategoryDetails
	if *v == nil {
		sv = &types.CategoryDetails{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "PointsOfInterest":
			if err := awsRestjson1_deserializeDocumentPointsOfInterest(&sv.PointsOfInterest, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentCharacterOffsets(v **types.CharacterOffsets, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.CharacterOffsets
	if *v == nil {
		sv = &types.CharacterOffsets{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "BeginOffsetChar":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected CharacterOffset to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.BeginOffsetChar = int32(i64)
			}

		case "EndOffsetChar":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected CharacterOffset to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.EndOffsetChar = int32(i64)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInternalServiceException(v **types.InternalServiceException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalServiceException
	if *v == nil {
		sv = &types.InternalServiceException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInvalidRequestException(v **types.InvalidRequestException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InvalidRequestException
	if *v == nil {
		sv = &types.InvalidRequestException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentIssueDetected(v **types.IssueDetected, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.IssueDetected
	if *v == nil {
		sv = &types.IssueDetected{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "CharacterOffsets":
			if err := awsRestjson1_deserializeDocumentCharacterOffsets(&sv.CharacterOffsets, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentIssuesDetected(v *[]types.IssueDetected, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.IssueDetected
	if *v == nil {
		cv = []types.IssueDetected{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.IssueDetected
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentIssueDetected(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentMatchedCategories(v *[]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []string
	if *v == nil {
		cv = []string{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected CategoryName to be of type string, got %T instead", value)
			}
			col = jtv
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentMatchedDetails(v *map[string]types.CategoryDetails, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var mv map[string]types.CategoryDetails
	if *v == nil {
		mv = map[string]types.CategoryDetails{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal types.CategoryDetails
		mapVar := parsedVal
		destAddr := &mapVar
		if err := awsRestjson1_deserializeDocumentCategoryDetails(&destAddr, value); err != nil {
			return err
		}
		parsedVal = *destAddr
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentPointOfInterest(v **types.PointOfInterest, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.PointOfInterest
	if *v == nil {
		sv = &types.PointOfInterest{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "BeginOffsetMillis":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected OffsetMillis to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.BeginOffsetMillis = int32(i64)
			}

		case "EndOffsetMillis":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected OffsetMillis to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.EndOffsetMillis = int32(i64)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentPointsOfInterest(v *[]types.PointOfInterest, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.PointOfInterest
	if *v == nil {
		cv = []types.PointOfInterest{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.PointOfInterest
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentPointOfInterest(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentRealtimeContactAnalysisSegment(v **types.RealtimeContactAnalysisSegment, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.RealtimeContactAnalysisSegment
	if *v == nil {
		sv = &types.RealtimeContactAnalysisSegment{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Categories":
			if err := awsRestjson1_deserializeDocumentCategories(&sv.Categories, value); err != nil {
				return err
			}

		case "Transcript":
			if err := awsRestjson1_deserializeDocumentTranscript(&sv.Transcript, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentRealtimeContactAnalysisSegments(v *[]types.RealtimeContactAnalysisSegment, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.RealtimeContactAnalysisSegment
	if *v == nil {
		cv = []types.RealtimeContactAnalysisSegment{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.RealtimeContactAnalysisSegment
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentRealtimeContactAnalysisSegment(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentResourceNotFoundException(v **types.ResourceNotFoundException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ResourceNotFoundException
	if *v == nil {
		sv = &types.ResourceNotFoundException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentThrottlingException(v **types.ThrottlingException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ThrottlingException
	if *v == nil {
		sv = &types.ThrottlingException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentTranscript(v **types.Transcript, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Transcript
	if *v == nil {
		sv = &types.Transcript{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "BeginOffsetMillis":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected OffsetMillis to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.BeginOffsetMillis = int32(i64)
			}

		case "Content":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected TranscriptContent to be of type string, got %T instead", value)
				}
				sv.Content = ptr.String(jtv)
			}

		case "EndOffsetMillis":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected OffsetMillis to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.EndOffsetMillis = int32(i64)
			}

		case "Id":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected TranscriptId to be of type string, got %T instead", value)
				}
				sv.Id = ptr.String(jtv)
			}

		case "IssuesDetected":
			if err := awsRestjson1_deserializeDocumentIssuesDetected(&sv.IssuesDetected, value); err != nil {
				return err
			}

		case "ParticipantId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ParticipantId to be of type string, got %T instead", value)
				}
				sv.ParticipantId = ptr.String(jtv)
			}

		case "ParticipantRole":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ParticipantRole to be of type string, got %T instead", value)
				}
				sv.ParticipantRole = ptr.String(jtv)
			}

		case "Sentiment":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected SentimentValue to be of type string, got %T instead", value)
				}
				sv.Sentiment = types.SentimentValue(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
