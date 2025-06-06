// Code generated by smithy-go-codegen DO NOT EDIT.

package invoicing

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/invoicing/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson10_serializeOpBatchGetInvoiceProfile struct {
}

func (*awsAwsjson10_serializeOpBatchGetInvoiceProfile) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpBatchGetInvoiceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchGetInvoiceProfileInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.BatchGetInvoiceProfile")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentBatchGetInvoiceProfileInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpCreateInvoiceUnit struct {
}

func (*awsAwsjson10_serializeOpCreateInvoiceUnit) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpCreateInvoiceUnit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateInvoiceUnitInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.CreateInvoiceUnit")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentCreateInvoiceUnitInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDeleteInvoiceUnit struct {
}

func (*awsAwsjson10_serializeOpDeleteInvoiceUnit) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDeleteInvoiceUnit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteInvoiceUnitInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.DeleteInvoiceUnit")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDeleteInvoiceUnitInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpGetInvoiceUnit struct {
}

func (*awsAwsjson10_serializeOpGetInvoiceUnit) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGetInvoiceUnit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetInvoiceUnitInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.GetInvoiceUnit")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGetInvoiceUnitInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListInvoiceSummaries struct {
}

func (*awsAwsjson10_serializeOpListInvoiceSummaries) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListInvoiceSummaries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListInvoiceSummariesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.ListInvoiceSummaries")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListInvoiceSummariesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListInvoiceUnits struct {
}

func (*awsAwsjson10_serializeOpListInvoiceUnits) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListInvoiceUnits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListInvoiceUnitsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.ListInvoiceUnits")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListInvoiceUnitsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListTagsForResource struct {
}

func (*awsAwsjson10_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.ListTagsForResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListTagsForResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpTagResource struct {
}

func (*awsAwsjson10_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.TagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUntagResource struct {
}

func (*awsAwsjson10_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.UntagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUntagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUpdateInvoiceUnit struct {
}

func (*awsAwsjson10_serializeOpUpdateInvoiceUnit) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUpdateInvoiceUnit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateInvoiceUnitInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Invoicing.UpdateInvoiceUnit")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUpdateInvoiceUnitInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsAwsjson10_serializeDocumentAccountIdList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10_serializeDocumentBillingPeriod(v *types.BillingPeriod, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Month != nil {
		ok := object.Key("Month")
		ok.Integer(*v.Month)
	}

	if v.Year != nil {
		ok := object.Key("Year")
		ok.Integer(*v.Year)
	}

	return nil
}

func awsAwsjson10_serializeDocumentDateInterval(v *types.DateInterval, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndDate != nil {
		ok := object.Key("EndDate")
		ok.Double(smithytime.FormatEpochSeconds(*v.EndDate))
	}

	if v.StartDate != nil {
		ok := object.Key("StartDate")
		ok.Double(smithytime.FormatEpochSeconds(*v.StartDate))
	}

	return nil
}

func awsAwsjson10_serializeDocumentFilters(v *types.Filters, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Accounts != nil {
		ok := object.Key("Accounts")
		if err := awsAwsjson10_serializeDocumentAccountIdList(v.Accounts, ok); err != nil {
			return err
		}
	}

	if v.InvoiceReceivers != nil {
		ok := object.Key("InvoiceReceivers")
		if err := awsAwsjson10_serializeDocumentAccountIdList(v.InvoiceReceivers, ok); err != nil {
			return err
		}
	}

	if v.Names != nil {
		ok := object.Key("Names")
		if err := awsAwsjson10_serializeDocumentInvoiceUnitNames(v.Names, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeDocumentInvoiceSummariesFilter(v *types.InvoiceSummariesFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BillingPeriod != nil {
		ok := object.Key("BillingPeriod")
		if err := awsAwsjson10_serializeDocumentBillingPeriod(v.BillingPeriod, ok); err != nil {
			return err
		}
	}

	if v.InvoicingEntity != nil {
		ok := object.Key("InvoicingEntity")
		ok.String(*v.InvoicingEntity)
	}

	if v.TimeInterval != nil {
		ok := object.Key("TimeInterval")
		if err := awsAwsjson10_serializeDocumentDateInterval(v.TimeInterval, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeDocumentInvoiceSummariesSelector(v *types.InvoiceSummariesSelector, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.ResourceType) > 0 {
		ok := object.Key("ResourceType")
		ok.String(string(v.ResourceType))
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson10_serializeDocumentInvoiceUnitNames(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10_serializeDocumentInvoiceUnitRule(v *types.InvoiceUnitRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.LinkedAccounts != nil {
		ok := object.Key("LinkedAccounts")
		if err := awsAwsjson10_serializeDocumentAccountIdList(v.LinkedAccounts, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeDocumentResourceTag(v *types.ResourceTag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson10_serializeDocumentResourceTagKeyList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10_serializeDocumentResourceTagList(v []types.ResourceTag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson10_serializeDocumentResourceTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10_serializeOpDocumentBatchGetInvoiceProfileInput(v *BatchGetInvoiceProfileInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AccountIds != nil {
		ok := object.Key("AccountIds")
		if err := awsAwsjson10_serializeDocumentAccountIdList(v.AccountIds, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentCreateInvoiceUnitInput(v *CreateInvoiceUnitInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Description != nil {
		ok := object.Key("Description")
		ok.String(*v.Description)
	}

	if v.InvoiceReceiver != nil {
		ok := object.Key("InvoiceReceiver")
		ok.String(*v.InvoiceReceiver)
	}

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.ResourceTags != nil {
		ok := object.Key("ResourceTags")
		if err := awsAwsjson10_serializeDocumentResourceTagList(v.ResourceTags, ok); err != nil {
			return err
		}
	}

	if v.Rule != nil {
		ok := object.Key("Rule")
		if err := awsAwsjson10_serializeDocumentInvoiceUnitRule(v.Rule, ok); err != nil {
			return err
		}
	}

	if v.TaxInheritanceDisabled {
		ok := object.Key("TaxInheritanceDisabled")
		ok.Boolean(v.TaxInheritanceDisabled)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDeleteInvoiceUnitInput(v *DeleteInvoiceUnitInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.InvoiceUnitArn != nil {
		ok := object.Key("InvoiceUnitArn")
		ok.String(*v.InvoiceUnitArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGetInvoiceUnitInput(v *GetInvoiceUnitInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AsOf != nil {
		ok := object.Key("AsOf")
		ok.Double(smithytime.FormatEpochSeconds(*v.AsOf))
	}

	if v.InvoiceUnitArn != nil {
		ok := object.Key("InvoiceUnitArn")
		ok.String(*v.InvoiceUnitArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListInvoiceSummariesInput(v *ListInvoiceSummariesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filter != nil {
		ok := object.Key("Filter")
		if err := awsAwsjson10_serializeDocumentInvoiceSummariesFilter(v.Filter, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.Selector != nil {
		ok := object.Key("Selector")
		if err := awsAwsjson10_serializeDocumentInvoiceSummariesSelector(v.Selector, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListInvoiceUnitsInput(v *ListInvoiceUnitsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AsOf != nil {
		ok := object.Key("AsOf")
		ok.Double(smithytime.FormatEpochSeconds(*v.AsOf))
	}

	if v.Filters != nil {
		ok := object.Key("Filters")
		if err := awsAwsjson10_serializeDocumentFilters(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListTagsForResourceInput(v *ListTagsForResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.ResourceTags != nil {
		ok := object.Key("ResourceTags")
		if err := awsAwsjson10_serializeDocumentResourceTagList(v.ResourceTags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUntagResourceInput(v *UntagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.ResourceTagKeys != nil {
		ok := object.Key("ResourceTagKeys")
		if err := awsAwsjson10_serializeDocumentResourceTagKeyList(v.ResourceTagKeys, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUpdateInvoiceUnitInput(v *UpdateInvoiceUnitInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Description != nil {
		ok := object.Key("Description")
		ok.String(*v.Description)
	}

	if v.InvoiceUnitArn != nil {
		ok := object.Key("InvoiceUnitArn")
		ok.String(*v.InvoiceUnitArn)
	}

	if v.Rule != nil {
		ok := object.Key("Rule")
		if err := awsAwsjson10_serializeDocumentInvoiceUnitRule(v.Rule, ok); err != nil {
			return err
		}
	}

	if v.TaxInheritanceDisabled != nil {
		ok := object.Key("TaxInheritanceDisabled")
		ok.Boolean(*v.TaxInheritanceDisabled)
	}

	return nil
}
