// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"bytes"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GreetingWithErrors_InvalidGreeting_awsAwsjson10Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.InvalidGreeting
	}{
		// Parses simple JSON errors
		"AwsJson10InvalidGreetingError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.json10#InvalidGreeting",
			    "Message": "Hi"
			}`),
			ExpectError: &types.InvalidGreeting{
				Message: ptr.String("Hi"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.InvalidGreeting operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.InvalidGreeting
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.InvalidGreeting result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_ComplexError_awsAwsjson10Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.ComplexError
	}{
		// Parses a complex error with no message member
		"AwsJson10ComplexError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.json10#ComplexError",
			    "TopLevel": "Top level",
			    "Nested": {
			        "Foo": "bar"
			    }
			}`),
			ExpectError: &types.ComplexError{
				TopLevel: ptr.String("Top level"),
				Nested: &types.ComplexNestedErrorData{
					Foo: ptr.String("bar"),
				},
			},
		},
		// Parses a complex error with an empty body
		"AwsJson10EmptyComplexError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.json10#ComplexError"
			}`),
			ExpectError: &types.ComplexError{},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.ComplexError operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.ComplexError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.ComplexError result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_FooError_awsAwsjson10Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.FooError
	}{
		// Serializes the X-Amzn-ErrorType header. For an example service, see Amazon EKS.
		"AwsJson10FooErrorUsingXAmznErrorType": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"FooError"},
			},
			ExpectError: &types.FooError{},
		},
		// Some X-Amzn-Errortype headers contain URLs. Clients need to split the URL on
		// ':' and take only the first half of the string. For example,
		// 'ValidationException:http://internal.amazon.com/coral/com.amazon.coral.validate/'
		// is to be interpreted as 'ValidationException'.
		//
		// For an example service see Amazon Polly.
		"AwsJson10FooErrorUsingXAmznErrorTypeWithUri": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"},
			},
			ExpectError: &types.FooError{},
		},
		// X-Amzn-Errortype might contain a URL and a namespace. Client should extract
		// only the shape name. This is a pathalogical case that might not actually happen
		// in any deployed AWS service.
		"AwsJson10FooErrorUsingXAmznErrorTypeWithUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"aws.protocoltests.json10#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"},
			},
			ExpectError: &types.FooError{},
		},
		// This example uses the 'code' property in the output rather than
		// X-Amzn-Errortype. Some services do this though it's preferable to send the
		// X-Amzn-Errortype. Client implementations must first check for the
		// X-Amzn-Errortype and then check for a top-level 'code' property.
		//
		// For example service see Amazon S3 Glacier.
		"AwsJson10FooErrorUsingCode": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using code, and it might contain a namespace.
		// Clients should just take the last part of the string after '#'.
		"AwsJson10FooErrorUsingCodeAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "aws.protocoltests.json10#FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using code, and it might contain a namespace. It
		// also might contain a URI. Clients should just take the last part of the string
		// after '#' and before ":". This is a pathalogical case that might not occur in
		// any deployed AWS service.
		"AwsJson10FooErrorUsingCodeUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "aws.protocoltests.json10#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type.
		"AwsJson10FooErrorWithDunderType": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type, and it might contain a namespace.
		// Clients should just take the last part of the string after '#'.
		"AwsJson10FooErrorWithDunderTypeAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.json10#FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type, and it might contain a namespace.
		// It also might contain a URI. Clients should just take the last part of the
		// string after '#' and before ":". This is a pathalogical case that might not
		// occur in any deployed AWS service.
		"AwsJson10FooErrorWithDunderTypeUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.json10#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"
			}`),
			ExpectError: &types.FooError{},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.FooError operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.FooError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.FooError result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}
