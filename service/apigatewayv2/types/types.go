// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Settings for logging access in a stage.
type AccessLogSettings struct {

	// The ARN of the CloudWatch Logs log group to receive access logs.
	DestinationArn *string

	// A single line format of the access logs of data, as specified by selected
	// $context variables. The format must include at least $context.requestId.
	Format *string

	noSmithyDocumentSerde
}

// Represents an API.
type Api struct {

	// The name of the API.
	//
	// This member is required.
	Name *string

	// The API protocol.
	//
	// This member is required.
	ProtocolType ProtocolType

	// The route selection expression for the API. For HTTP APIs, the
	// routeSelectionExpression must be ${request.method} ${request.path}. If not
	// provided, this will be the default for HTTP APIs. This property is required for
	// WebSocket APIs.
	//
	// This member is required.
	RouteSelectionExpression *string

	// The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com.
	// The stage name is typically appended to this URI to form a complete path to a
	// deployed API stage.
	ApiEndpoint *string

	// Specifies whether an API is managed by API Gateway. You can't update or delete
	// a managed API by using API Gateway. A managed API can be deleted only through
	// the tooling or service that created it.
	ApiGatewayManaged *bool

	// The API ID.
	ApiId *string

	// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions].
	//
	// [API Key Selection Expressions]: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions
	ApiKeySelectionExpression *string

	// A CORS configuration. Supported only for HTTP APIs.
	CorsConfiguration *Cors

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The description of the API.
	Description *string

	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint *bool

	// Avoid validating models when creating a deployment. Supported only for
	// WebSocket APIs.
	DisableSchemaValidation *bool

	// The validation information during API import. This may include particular
	// properties of your OpenAPI definition which are ignored during import. Supported
	// only for HTTP APIs.
	ImportInfo []string

	// The IP address types that can invoke the API.
	IpAddressType IpAddressType

	// A collection of tags associated with the API.
	Tags map[string]string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string

	noSmithyDocumentSerde
}

// Represents an API mapping.
type ApiMapping struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The API stage.
	//
	// This member is required.
	Stage *string

	// The API mapping identifier.
	ApiMappingId *string

	// The API mapping key.
	ApiMappingKey *string

	noSmithyDocumentSerde
}

// Represents an authorizer.
type Authorizer struct {

	// The name of the authorizer.
	//
	// This member is required.
	Name *string

	// Specifies the required credentials as an IAM role for API Gateway to invoke the
	// authorizer. To specify an IAM role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To use resource-based permissions on the Lambda
	// function, don't specify this parameter. Supported only for REQUEST authorizers.
	AuthorizerCredentialsArn *string

	// The authorizer identifier.
	AuthorizerId *string

	// Specifies the format of the payload sent to an HTTP API Lambda authorizer.
	// Required for HTTP API Lambda authorizers. Supported values are 1.0 and 2.0. To
	// learn more, see [Working with AWS Lambda authorizers for HTTP APIs].
	//
	// [Working with AWS Lambda authorizers for HTTP APIs]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html
	AuthorizerPayloadFormatVersion *string

	// The time to live (TTL) for cached authorizer results, in seconds. If it equals
	// 0, authorization caching is disabled. If it is greater than 0, API Gateway
	// caches authorizer responses. The maximum value is 3600, or 1 hour. Supported
	// only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds *int32

	// The authorizer type. Specify REQUEST for a Lambda function using incoming
	// request parameters. Specify JWT to use JSON Web Tokens (supported only for HTTP
	// APIs).
	AuthorizerType AuthorizerType

	// The authorizer's Uniform Resource Identifier (URI). For REQUEST authorizers,
	// this must be a well-formed Lambda function URI, for example,
	// arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form:
	// arn:aws:apigateway:{region}:lambda:path/{service_api} , where {region} is the
	// same as the region hosting the Lambda function, path indicates that the
	// remaining substring in the URI should be treated as the path to the resource,
	// including the initial /. For Lambda functions, this is usually of the form
	// /2015-03-31/functions/[FunctionARN]/invocations. Supported only for REQUEST
	// authorizers.
	AuthorizerUri *string

	// Specifies whether a Lambda authorizer returns a response in a simple format. If
	// enabled, the Lambda authorizer can return a boolean value instead of an IAM
	// policy. Supported only for HTTP APIs. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs]
	//
	// [Working with AWS Lambda authorizers for HTTP APIs]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html
	EnableSimpleResponses *bool

	// The identity source for which authorization is requested.
	//
	// For a REQUEST authorizer, this is optional. The value is a set of one or more
	// mapping expressions of the specified request parameters. The identity source can
	// be headers, query string parameters, stage variables, and context parameters.
	// For example, if an Auth header and a Name query string parameter are defined as
	// identity sources, this value is route.request.header.Auth,
	// route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection
	// expressions prefixed with $, for example, $request.header.Auth,
	// $request.querystring.Name. These parameters are used to perform runtime
	// validation for Lambda-based authorizers by verifying all of the identity-related
	// request parameters are present in the request, not null, and non-empty. Only
	// when this is true does the authorizer invoke the authorizer Lambda function.
	// Otherwise, it returns a 401 Unauthorized response without calling the Lambda
	// function. For HTTP APIs, identity sources are also used as the cache key when
	// caching is enabled. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs].
	//
	// For JWT, a single entry that specifies where to extract the JSON Web Token
	// (JWT) from inbound requests. Currently only header-based and query
	// parameter-based selections are supported, for example
	// $request.header.Authorization.
	//
	// [Working with AWS Lambda authorizers for HTTP APIs]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html
	IdentitySource []string

	// The validation expression does not apply to the REQUEST authorizer.
	IdentityValidationExpression *string

	// Represents the configuration of a JWT authorizer. Required for the JWT
	// authorizer type. Supported only for HTTP APIs.
	JwtConfiguration *JWTConfiguration

	noSmithyDocumentSerde
}

// Represents a CORS configuration. Supported only for HTTP APIs. See [Configuring CORS] for more
// information.
//
// [Configuring CORS]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html
type Cors struct {

	// Specifies whether credentials are included in the CORS request. Supported only
	// for HTTP APIs.
	AllowCredentials *bool

	// Represents a collection of allowed headers. Supported only for HTTP APIs.
	AllowHeaders []string

	// Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.
	AllowMethods []string

	// Represents a collection of allowed origins. Supported only for HTTP APIs.
	AllowOrigins []string

	// Represents a collection of exposed headers. Supported only for HTTP APIs.
	ExposeHeaders []string

	// The number of seconds that the browser should cache preflight request results.
	// Supported only for HTTP APIs.
	MaxAge *int32

	noSmithyDocumentSerde
}

// An immutable representation of an API that can be called by users. A Deployment
// must be associated with a Stage for it to be callable over the internet.
type Deployment struct {

	// Specifies whether a deployment was automatically released.
	AutoDeployed *bool

	// The date and time when the Deployment resource was created.
	CreatedDate *time.Time

	// The identifier for the deployment.
	DeploymentId *string

	// The status of the deployment: PENDING, FAILED, or SUCCEEDED.
	DeploymentStatus DeploymentStatus

	// May contain additional feedback on the status of an API deployment.
	DeploymentStatusMessage *string

	// The description for the deployment.
	Description *string

	noSmithyDocumentSerde
}

// Represents a domain name.
type DomainName struct {

	// The name of the DomainName resource.
	//
	// This member is required.
	DomainName *string

	// The API mapping selection expression.
	ApiMappingSelectionExpression *string

	// Represents an Amazon Resource Name (ARN).
	DomainNameArn *string

	// The domain name configurations.
	DomainNameConfigurations []DomainNameConfiguration

	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication *MutualTlsAuthentication

	// The routing mode.
	RoutingMode RoutingMode

	// The collection of tags associated with a domain name.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The domain name configuration.
type DomainNameConfiguration struct {

	// A domain name for the API.
	ApiGatewayDomainName *string

	// An AWS-managed certificate that will be used by the edge-optimized endpoint for
	// this domain name. AWS Certificate Manager is the only supported source.
	CertificateArn *string

	// The user-friendly name of the certificate that will be used by the
	// edge-optimized endpoint for this domain name.
	CertificateName *string

	// The timestamp when the certificate that was used by edge-optimized endpoint for
	// this domain name was uploaded.
	CertificateUploadDate *time.Time

	// The status of the domain name migration. The valid values are AVAILABLE,
	// UPDATING, PENDING_CERTIFICATE_REIMPORT, and PENDING_OWNERSHIP_VERIFICATION. If
	// the status is UPDATING, the domain cannot be modified further until the existing
	// operation is complete. If it is AVAILABLE, the domain can be updated.
	DomainNameStatus DomainNameStatus

	// An optional text message containing detailed information about status of the
	// domain name migration.
	DomainNameStatusMessage *string

	// The endpoint type.
	EndpointType EndpointType

	// The Amazon Route 53 Hosted Zone ID of the endpoint.
	HostedZoneId *string

	// The IP address types that can invoke the domain name. Use ipv4 to allow only
	// IPv4 addresses to invoke your domain name, or use dualstack to allow both IPv4
	// and IPv6 addresses to invoke your domain name.
	IpAddressType IpAddressType

	// The ARN of the public certificate issued by ACM to validate ownership of your
	// custom domain. Only required when configuring mutual TLS and using an ACM
	// imported or private CA certificate ARN as the regionalCertificateArn
	OwnershipVerificationCertificateArn *string

	// The Transport Layer Security (TLS) version of the security policy for this
	// domain name. The valid values are TLS_1_0 and TLS_1_2.
	SecurityPolicy SecurityPolicy

	noSmithyDocumentSerde
}

// Represents an integration.
type Integration struct {

	// Specifies whether an integration is managed by API Gateway. If you created an
	// API using using quick create, the resulting integration is managed by API
	// Gateway. You can update a managed integration, but you can't delete it.
	ApiGatewayManaged *bool

	// The ID of the VPC link for a private integration. Supported only for HTTP APIs.
	ConnectionId *string

	// The type of the network connection to the integration endpoint. Specify
	// INTERNET for connections through the public routable internet or VPC_LINK for
	// private connections between API Gateway and resources in a VPC. The default
	// value is INTERNET.
	ConnectionType ConnectionType

	// Supported only for WebSocket APIs. Specifies how to handle response payload
	// content type conversions. Supported values are CONVERT_TO_BINARY and
	// CONVERT_TO_TEXT, with the following behaviors:
	//
	// CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string to
	// the corresponding binary blob.
	//
	// CONVERT_TO_TEXT: Converts a response payload from a binary blob to a
	// Base64-encoded string.
	//
	// If this property is not defined, the response payload will be passed through
	// from the integration response to the route response or method response without
	// modification.
	ContentHandlingStrategy ContentHandlingStrategy

	// Specifies the credentials required for the integration, if any. For AWS
	// integrations, three options are available. To specify an IAM Role for API
	// Gateway to assume, use the role's Amazon Resource Name (ARN). To require that
	// the caller's identity be passed through from the request, specify the string
	// arn:aws:iam::*:user/*. To use resource-based permissions on supported AWS
	// services, specify null.
	CredentialsArn *string

	// Represents the description of an integration.
	Description *string

	// Represents the identifier of an integration.
	IntegrationId *string

	// Specifies the integration's HTTP method type.
	IntegrationMethod *string

	// The integration response selection expression for the integration. Supported
	// only for WebSocket APIs. See [Integration Response Selection Expressions].
	//
	// [Integration Response Selection Expressions]: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions
	IntegrationResponseSelectionExpression *string

	// Supported only for HTTP API AWS_PROXY integrations. Specifies the AWS service
	// action to invoke. To learn more, see [Integration subtype reference].
	//
	// [Integration subtype reference]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html
	IntegrationSubtype *string

	// The integration type of an integration. One of the following:
	//
	// AWS: for integrating the route or method request with an AWS service action,
	// including the Lambda function-invoking action. With the Lambda function-invoking
	// action, this is referred to as the Lambda custom integration. With any other AWS
	// service action, this is known as AWS integration. Supported only for WebSocket
	// APIs.
	//
	// AWS_PROXY: for integrating the route or method request with a Lambda function
	// or other AWS service action. This integration is also referred to as a Lambda
	// proxy integration.
	//
	// HTTP: for integrating the route or method request with an HTTP endpoint. This
	// integration is also referred to as the HTTP custom integration. Supported only
	// for WebSocket APIs.
	//
	// HTTP_PROXY: for integrating the route or method request with an HTTP endpoint,
	// with the client request passed through as-is. This is also referred to as HTTP
	// proxy integration.
	//
	// MOCK: for integrating the route or method request with API Gateway as a
	// "loopback" endpoint without invoking any backend. Supported only for WebSocket
	// APIs.
	IntegrationType IntegrationType

	// For a Lambda integration, specify the URI of a Lambda function.
	//
	// For an HTTP integration, specify a fully-qualified URL.
	//
	// For an HTTP API private integration, specify the ARN of an Application Load
	// Balancer listener, Network Load Balancer listener, or AWS Cloud Map service. If
	// you specify the ARN of an AWS Cloud Map service, API Gateway uses
	// DiscoverInstances to identify resources. You can use query parameters to target
	// specific resources. To learn more, see [DiscoverInstances]. For private integrations, all
	// resources must be owned by the same AWS account.
	//
	// [DiscoverInstances]: https://docs.aws.amazon.com/cloud-map/latest/api/API_DiscoverInstances.html
	IntegrationUri *string

	// Specifies the pass-through behavior for incoming requests based on the
	// Content-Type header in the request, and the available mapping templates
	// specified as the requestTemplates property on the Integration resource. There
	// are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and NEVER. Supported
	// only for WebSocket APIs.
	//
	// WHEN_NO_MATCH passes the request body for unmapped content types through to the
	// integration backend without transformation.
	//
	// NEVER rejects unmapped content types with an HTTP 415 Unsupported Media Type
	// response.
	//
	// WHEN_NO_TEMPLATES allows pass-through when the integration has no content types
	// mapped to templates. However, if there is at least one content type defined,
	// unmapped content types will be rejected with the same HTTP 415 Unsupported Media
	// Type response.
	PassthroughBehavior PassthroughBehavior

	// Specifies the format of the payload sent to an integration. Required for HTTP
	// APIs. Supported values for Lambda proxy integrations are 1.0 and 2.0. For all
	// other integrations, 1.0 is the only supported value. To learn more, see [Working with AWS Lambda proxy integrations for HTTP APIs].
	//
	// [Working with AWS Lambda proxy integrations for HTTP APIs]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	PayloadFormatVersion *string

	// For WebSocket APIs, a key-value map specifying request parameters that are
	// passed from the method request to the backend. The key is an integration request
	// parameter name and the associated value is a method request parameter value or
	// static value that must be enclosed within single quotes and pre-encoded as
	// required by the backend. The method request parameter value must match the
	// pattern of method.request.{location}.{name} , where {location} is querystring,
	// path, or header; and {name} must be a valid and unique method request parameter
	// name.
	//
	// For HTTP API integrations with a specified integrationSubtype, request
	// parameters are a key-value map specifying parameters that are passed to
	// AWS_PROXY integrations. You can provide static values, or map request data,
	// stage variables, or context variables that are evaluated at runtime. To learn
	// more, see [Working with AWS service integrations for HTTP APIs].
	//
	// For HTTP API integrations, without a specified integrationSubtype request
	// parameters are a key-value map specifying how to transform HTTP requests before
	// sending them to backend integrations. The key should follow the pattern
	// <action>:<header|querystring|path>.<location>. The action can be append,
	// overwrite or remove. For values, you can provide static values, or map request
	// data, stage variables, or context variables that are evaluated at runtime. To
	// learn more, see [Transforming API requests and responses].
	//
	// [Working with AWS service integrations for HTTP APIs]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services.html
	// [Transforming API requests and responses]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	RequestParameters map[string]string

	// Represents a map of Velocity templates that are applied on the request payload
	// based on the value of the Content-Type header sent by the client. The content
	// type value is the key in this map, and the template (as a String) is the value.
	// Supported only for WebSocket APIs.
	RequestTemplates map[string]string

	// Supported only for HTTP APIs. You use response parameters to transform the HTTP
	// response from a backend integration before returning the response to clients.
	// Specify a key-value map from a selection key to response parameters. The
	// selection key must be a valid HTTP status code within the range of 200-599.
	// Response parameters are a key-value map. The key must match pattern
	// <action>:<header>.<location> or overwrite.statuscode. The action can be append,
	// overwrite or remove. The value can be a static value, or map to response data,
	// stage variables, or context variables that are evaluated at runtime. To learn
	// more, see [Transforming API requests and responses].
	//
	// [Transforming API requests and responses]: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	ResponseParameters map[string]map[string]string

	// The template selection expression for the integration. Supported only for
	// WebSocket APIs.
	TemplateSelectionExpression *string

	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and
	// between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29
	// seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	TimeoutInMillis *int32

	// The TLS configuration for a private integration. If you specify a TLS
	// configuration, private integration traffic uses the HTTPS protocol. Supported
	// only for HTTP APIs.
	TlsConfig *TlsConfig

	noSmithyDocumentSerde
}

// Represents an integration response.
type IntegrationResponse struct {

	// The integration response key.
	//
	// This member is required.
	IntegrationResponseKey *string

	// Supported only for WebSocket APIs. Specifies how to handle response payload
	// content type conversions. Supported values are CONVERT_TO_BINARY and
	// CONVERT_TO_TEXT, with the following behaviors:
	//
	// CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string to
	// the corresponding binary blob.
	//
	// CONVERT_TO_TEXT: Converts a response payload from a binary blob to a
	// Base64-encoded string.
	//
	// If this property is not defined, the response payload will be passed through
	// from the integration response to the route response or method response without
	// modification.
	ContentHandlingStrategy ContentHandlingStrategy

	// The integration response ID.
	IntegrationResponseId *string

	// A key-value map specifying response parameters that are passed to the method
	// response from the backend. The key is a method response header parameter name
	// and the mapped value is an integration response header value, a static value
	// enclosed within a pair of single quotes, or a JSON expression from the
	// integration response body. The mapping key must match the pattern of
	// method.response.header.{name}, where name is a valid and unique header name. The
	// mapped non-static value must match the pattern of
	// integration.response.header.{name} or
	// integration.response.body.{JSON-expression}, where name is a valid and unique
	// response header name and JSON-expression is a valid JSON expression without the
	// $ prefix.
	ResponseParameters map[string]string

	// The collection of response templates for the integration response as a
	// string-to-string map of key-value pairs. Response templates are represented as a
	// key/value map, with a content-type as the key and a template as the value.
	ResponseTemplates map[string]string

	// The template selection expressions for the integration response.
	TemplateSelectionExpression *string

	noSmithyDocumentSerde
}

// Represents the configuration of a JWT authorizer. Required for the JWT
// authorizer type. Supported only for HTTP APIs.
type JWTConfiguration struct {

	// A list of the intended recipients of the JWT. A valid JWT must provide an aud
	// that matches at least one entry in this list. See [RFC 7519]. Supported only for HTTP
	// APIs.
	//
	// [RFC 7519]: https://tools.ietf.org/html/rfc7519#section-4.1.3
	Audience []string

	// The base domain of the identity provider that issues JSON Web Tokens. For
	// example, an Amazon Cognito user pool has the following format:
	// https://cognito-idp.{region}.amazonaws.com/{userPoolId} . Required for the JWT
	// authorizer type. Supported only for HTTP APIs.
	Issuer *string

	noSmithyDocumentSerde
}

// Represents a data model for an API. Supported only for WebSocket APIs. See [Create Models and Mapping Templates for Request and Response Mappings].
//
// [Create Models and Mapping Templates for Request and Response Mappings]: https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html
type Model struct {

	// The name of the model. Must be alphanumeric.
	//
	// This member is required.
	Name *string

	// The content-type for the model, for example, "application/json".
	ContentType *string

	// The description of the model.
	Description *string

	// The model identifier.
	ModelId *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 model.
	Schema *string

	noSmithyDocumentSerde
}

type MutualTlsAuthentication struct {

	// An Amazon S3 URL that specifies the truststore for mutual TLS authentication,
	// for example, s3://bucket-name/key-name. The truststore can contain certificates
	// from public or private certificate authorities. To update the truststore, upload
	// a new version to S3, and then update your custom domain name to use the new
	// version. To update the truststore, you must have permissions to access the S3
	// object.
	TruststoreUri *string

	// The version of the S3 object that contains your truststore. To specify a
	// version, you must have versioning enabled for the S3 bucket.
	TruststoreVersion *string

	// A list of warnings that API Gateway returns while processing your truststore.
	// Invalid certificates produce warnings. Mutual TLS is still enabled, but some
	// clients might not be able to access your API. To resolve warnings, upload a new
	// truststore to S3, and then update you domain name to use the new version.
	TruststoreWarnings []string

	noSmithyDocumentSerde
}

type MutualTlsAuthenticationInput struct {

	// An Amazon S3 URL that specifies the truststore for mutual TLS authentication,
	// for example, s3://bucket-name/key-name. The truststore can contain certificates
	// from public or private certificate authorities. To update the truststore, upload
	// a new version to S3, and then update your custom domain name to use the new
	// version. To update the truststore, you must have permissions to access the S3
	// object.
	TruststoreUri *string

	// The version of the S3 object that contains your truststore. To specify a
	// version, you must have versioning enabled for the S3 bucket.
	TruststoreVersion *string

	noSmithyDocumentSerde
}

// Validation constraints imposed on parameters of a request (path, query string,
// headers).
type ParameterConstraints struct {

	// Whether or not the parameter is required.
	Required *bool

	noSmithyDocumentSerde
}

// Represents a route.
type Route struct {

	// The route key for the route.
	//
	// This member is required.
	RouteKey *string

	// Specifies whether a route is managed by API Gateway. If you created an API
	// using quick create, the $default route is managed by API Gateway. You can't
	// modify the $default route key.
	ApiGatewayManaged *bool

	// Specifies whether an API key is required for this route. Supported only for
	// WebSocket APIs.
	ApiKeyRequired *bool

	// A list of authorization scopes configured on a route. The scopes are used with
	// a JWT authorizer to authorize the method invocation. The authorization works by
	// matching the route scopes against the scopes parsed from the access token in the
	// incoming request. The method invocation is authorized if any route scope matches
	// a claimed scope in the access token. Otherwise, the invocation is not
	// authorized. When the route scope is configured, the client must provide an
	// access token instead of an identity token for authorization purposes.
	AuthorizationScopes []string

	// The authorization type for the route. For WebSocket APIs, valid values are NONE
	// for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a
	// Lambda authorizer For HTTP APIs, valid values are NONE for open access, JWT for
	// using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM for
	// using a Lambda authorizer.
	AuthorizationType AuthorizationType

	// The identifier of the Authorizer resource to be associated with this route. The
	// authorizer identifier is generated by API Gateway when you created the
	// authorizer.
	AuthorizerId *string

	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression *string

	// The operation name for the route.
	OperationName *string

	// The request models for the route. Supported only for WebSocket APIs.
	RequestModels map[string]string

	// The request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters map[string]ParameterConstraints

	// The route ID.
	RouteId *string

	// The route response selection expression for the route. Supported only for
	// WebSocket APIs.
	RouteResponseSelectionExpression *string

	// The target for the route.
	Target *string

	noSmithyDocumentSerde
}

// Represents a route response.
type RouteResponse struct {

	// Represents the route response key of a route response.
	//
	// This member is required.
	RouteResponseKey *string

	// Represents the model selection expression of a route response. Supported only
	// for WebSocket APIs.
	ModelSelectionExpression *string

	// Represents the response models of a route response.
	ResponseModels map[string]string

	// Represents the response parameters of a route response.
	ResponseParameters map[string]ParameterConstraints

	// Represents the identifier of a route response.
	RouteResponseId *string

	noSmithyDocumentSerde
}

// Represents a collection of route settings.
type RouteSettings struct {

	// Specifies whether (true) or not (false) data trace logging is enabled for this
	// route. This property affects the log entries pushed to Amazon CloudWatch Logs.
	// Supported only for WebSocket APIs.
	DataTraceEnabled *bool

	// Specifies whether detailed metrics are enabled.
	DetailedMetricsEnabled *bool

	// Specifies the logging level for this route: INFO, ERROR, or OFF. This property
	// affects the log entries pushed to Amazon CloudWatch Logs. Supported only for
	// WebSocket APIs.
	LoggingLevel LoggingLevel

	// Specifies the throttling burst limit.
	ThrottlingBurstLimit *int32

	// Specifies the throttling rate limit.
	ThrottlingRateLimit *float64

	noSmithyDocumentSerde
}

// Represents a routing rule.
type RoutingRule struct {

	// The routing rule action.
	Actions []RoutingRuleAction

	// The routing rule condition.
	Conditions []RoutingRuleCondition

	// The routing rule priority.
	Priority *int32

	// The routing rule ARN.
	RoutingRuleArn *string

	// The routing rule ID.
	RoutingRuleId *string

	noSmithyDocumentSerde
}

// The routing rule action.
type RoutingRuleAction struct {

	// Represents an InvokeApi action.
	//
	// This member is required.
	InvokeApi *RoutingRuleActionInvokeApi

	noSmithyDocumentSerde
}

// Represents an InvokeApi action.
type RoutingRuleActionInvokeApi struct {

	// The identifier.
	//
	// This member is required.
	ApiId *string

	// A string with a length between [1-128].
	//
	// This member is required.
	Stage *string

	// The strip base path setting.
	StripBasePath *bool

	noSmithyDocumentSerde
}

// Represents a routing rule condition.
type RoutingRuleCondition struct {

	// The base path to be matched.
	MatchBasePaths *RoutingRuleMatchBasePaths

	// The headers to be matched.
	MatchHeaders *RoutingRuleMatchHeaders

	noSmithyDocumentSerde
}

// Represents a MatchBasePaths condition.
type RoutingRuleMatchBasePaths struct {

	// The string of the case sensitive base path to be matched.
	//
	// This member is required.
	AnyOf []string

	noSmithyDocumentSerde
}

// Represents a MatchHeaders condition.
type RoutingRuleMatchHeaders struct {

	// The header name and header value glob to be matched. The matchHeaders condition
	// is matched if any of the header name and header value globs are matched.
	//
	// This member is required.
	AnyOf []RoutingRuleMatchHeaderValue

	noSmithyDocumentSerde
}

// Represents a MatchHeaderValue.
type RoutingRuleMatchHeaderValue struct {

	// After evaluating a selection expression, the result is compared against one or
	// more selection keys to find a matching key. See [Selection Expressions]for a list of expressions and
	// each expression's associated selection key type.
	//
	// [Selection Expressions]: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions
	//
	// This member is required.
	Header *string

	// An expression used to extract information at runtime. See [Selection Expressions] for more information.
	//
	// [Selection Expressions]: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions
	//
	// This member is required.
	ValueGlob *string

	noSmithyDocumentSerde
}

// Represents an API stage.
type Stage struct {

	// The name of the stage.
	//
	// This member is required.
	StageName *string

	// Settings for logging access in this stage.
	AccessLogSettings *AccessLogSettings

	// Specifies whether a stage is managed by API Gateway. If you created an API
	// using quick create, the $default stage is managed by API Gateway. You can't
	// modify the $default stage.
	ApiGatewayManaged *bool

	// Specifies whether updates to an API automatically trigger a new deployment. The
	// default value is false.
	AutoDeploy *bool

	// The identifier of a client certificate for a Stage. Supported only for
	// WebSocket APIs.
	ClientCertificateId *string

	// The timestamp when the stage was created.
	CreatedDate *time.Time

	// Default route settings for the stage.
	DefaultRouteSettings *RouteSettings

	// The identifier of the Deployment that the Stage is associated with. Can't be
	// updated if autoDeploy is enabled.
	DeploymentId *string

	// The description of the stage.
	Description *string

	// Describes the status of the last deployment of a stage. Supported only for
	// stages with autoDeploy enabled.
	LastDeploymentStatusMessage *string

	// The timestamp when the stage was last updated.
	LastUpdatedDate *time.Time

	// Route settings for the stage, by routeKey.
	RouteSettings map[string]RouteSettings

	// A map that defines the stage variables for a stage resource. Variable names can
	// have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	StageVariables map[string]string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The TLS configuration for a private integration. If you specify a TLS
// configuration, private integration traffic uses the HTTPS protocol. Supported
// only for HTTP APIs.
type TlsConfig struct {

	// If you specify a server name, API Gateway uses it to verify the hostname on the
	// integration's certificate. The server name is also included in the TLS handshake
	// to support Server Name Indication (SNI) or virtual hosting.
	ServerNameToVerify *string

	noSmithyDocumentSerde
}

// The TLS configuration for a private integration. If you specify a TLS
// configuration, private integration traffic uses the HTTPS protocol. Supported
// only for HTTP APIs.
type TlsConfigInput struct {

	// If you specify a server name, API Gateway uses it to verify the hostname on the
	// integration's certificate. The server name is also included in the TLS handshake
	// to support Server Name Indication (SNI) or virtual hosting.
	ServerNameToVerify *string

	noSmithyDocumentSerde
}

// Represents a VPC link.
type VpcLink struct {

	// The name of the VPC link.
	//
	// This member is required.
	Name *string

	// A list of security group IDs for the VPC link.
	//
	// This member is required.
	SecurityGroupIds []string

	// A list of subnet IDs to include in the VPC link.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the VPC link.
	//
	// This member is required.
	VpcLinkId *string

	// The timestamp when the VPC link was created.
	CreatedDate *time.Time

	// Tags for the VPC link.
	Tags map[string]string

	// The status of the VPC link.
	VpcLinkStatus VpcLinkStatus

	// A message summarizing the cause of the status of the VPC link.
	VpcLinkStatusMessage *string

	// The version of the VPC link.
	VpcLinkVersion VpcLinkVersion

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
