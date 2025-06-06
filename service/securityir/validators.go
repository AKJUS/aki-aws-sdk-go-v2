// Code generated by smithy-go-codegen DO NOT EDIT.

package securityir

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/securityir/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchGetMemberAccountDetails struct {
}

func (*validateOpBatchGetMemberAccountDetails) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetMemberAccountDetails) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetMemberAccountDetailsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetMemberAccountDetailsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelMembership struct {
}

func (*validateOpCancelMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCloseCase struct {
}

func (*validateOpCloseCase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCloseCase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CloseCaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCloseCaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCaseComment struct {
}

func (*validateOpCreateCaseComment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCaseComment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCaseCommentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCaseCommentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCase struct {
}

func (*validateOpCreateCase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMembership struct {
}

func (*validateOpCreateMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCaseAttachmentDownloadUrl struct {
}

func (*validateOpGetCaseAttachmentDownloadUrl) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCaseAttachmentDownloadUrl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCaseAttachmentDownloadUrlInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCaseAttachmentDownloadUrlInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCaseAttachmentUploadUrl struct {
}

func (*validateOpGetCaseAttachmentUploadUrl) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCaseAttachmentUploadUrl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCaseAttachmentUploadUrlInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCaseAttachmentUploadUrlInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCase struct {
}

func (*validateOpGetCase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMembership struct {
}

func (*validateOpGetMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListCaseEdits struct {
}

func (*validateOpListCaseEdits) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListCaseEdits) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListCaseEditsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListCaseEditsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListComments struct {
}

func (*validateOpListComments) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListComments) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListCommentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListCommentsInput(input); err != nil {
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

type validateOpUpdateCaseComment struct {
}

func (*validateOpUpdateCaseComment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCaseComment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCaseCommentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCaseCommentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCase struct {
}

func (*validateOpUpdateCase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCaseStatus struct {
}

func (*validateOpUpdateCaseStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCaseStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCaseStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCaseStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateMembership struct {
}

func (*validateOpUpdateMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateResolverType struct {
}

func (*validateOpUpdateResolverType) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateResolverType) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateResolverTypeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateResolverTypeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchGetMemberAccountDetailsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetMemberAccountDetails{}, middleware.After)
}

func addOpCancelMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelMembership{}, middleware.After)
}

func addOpCloseCaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCloseCase{}, middleware.After)
}

func addOpCreateCaseCommentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCaseComment{}, middleware.After)
}

func addOpCreateCaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCase{}, middleware.After)
}

func addOpCreateMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMembership{}, middleware.After)
}

func addOpGetCaseAttachmentDownloadUrlValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCaseAttachmentDownloadUrl{}, middleware.After)
}

func addOpGetCaseAttachmentUploadUrlValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCaseAttachmentUploadUrl{}, middleware.After)
}

func addOpGetCaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCase{}, middleware.After)
}

func addOpGetMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMembership{}, middleware.After)
}

func addOpListCaseEditsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListCaseEdits{}, middleware.After)
}

func addOpListCommentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListComments{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateCaseCommentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCaseComment{}, middleware.After)
}

func addOpUpdateCaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCase{}, middleware.After)
}

func addOpUpdateCaseStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCaseStatus{}, middleware.After)
}

func addOpUpdateMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMembership{}, middleware.After)
}

func addOpUpdateResolverTypeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateResolverType{}, middleware.After)
}

func validateImpactedAwsRegion(v *types.ImpactedAwsRegion) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImpactedAwsRegion"}
	if len(v.Region) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Region"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateImpactedAwsRegionList(v []types.ImpactedAwsRegion) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImpactedAwsRegionList"}
	for i := range v {
		if err := validateImpactedAwsRegion(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIncidentResponder(v *types.IncidentResponder) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IncidentResponder"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.JobTitle == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobTitle"))
	}
	if v.Email == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Email"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIncidentResponseTeam(v []types.IncidentResponder) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IncidentResponseTeam"}
	for i := range v {
		if err := validateIncidentResponder(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOptInFeature(v *types.OptInFeature) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OptInFeature"}
	if len(v.FeatureName) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureName"))
	}
	if v.IsEnabled == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IsEnabled"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOptInFeatures(v []types.OptInFeature) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OptInFeatures"}
	for i := range v {
		if err := validateOptInFeature(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateThreatActorIp(v *types.ThreatActorIp) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ThreatActorIp"}
	if v.IpAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddress"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateThreatActorIpList(v []types.ThreatActorIp) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ThreatActorIpList"}
	for i := range v {
		if err := validateThreatActorIp(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateWatcher(v *types.Watcher) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Watcher"}
	if v.Email == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Email"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateWatchers(v []types.Watcher) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Watchers"}
	for i := range v {
		if err := validateWatcher(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetMemberAccountDetailsInput(v *BatchGetMemberAccountDetailsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetMemberAccountDetailsInput"}
	if v.MembershipId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MembershipId"))
	}
	if v.AccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelMembershipInput(v *CancelMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelMembershipInput"}
	if v.MembershipId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MembershipId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCloseCaseInput(v *CloseCaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CloseCaseInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCaseCommentInput(v *CreateCaseCommentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCaseCommentInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if v.Body == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Body"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCaseInput(v *CreateCaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCaseInput"}
	if len(v.ResolverType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverType"))
	}
	if v.Title == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Title"))
	}
	if v.Description == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Description"))
	}
	if len(v.EngagementType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("EngagementType"))
	}
	if v.ReportedIncidentStartDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportedIncidentStartDate"))
	}
	if v.ImpactedAccounts == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ImpactedAccounts"))
	}
	if v.Watchers == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Watchers"))
	} else if v.Watchers != nil {
		if err := validateWatchers(v.Watchers); err != nil {
			invalidParams.AddNested("Watchers", err.(smithy.InvalidParamsError))
		}
	}
	if v.ThreatActorIpAddresses != nil {
		if err := validateThreatActorIpList(v.ThreatActorIpAddresses); err != nil {
			invalidParams.AddNested("ThreatActorIpAddresses", err.(smithy.InvalidParamsError))
		}
	}
	if v.ImpactedAwsRegions != nil {
		if err := validateImpactedAwsRegionList(v.ImpactedAwsRegions); err != nil {
			invalidParams.AddNested("ImpactedAwsRegions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMembershipInput(v *CreateMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMembershipInput"}
	if v.MembershipName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MembershipName"))
	}
	if v.IncidentResponseTeam == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IncidentResponseTeam"))
	} else if v.IncidentResponseTeam != nil {
		if err := validateIncidentResponseTeam(v.IncidentResponseTeam); err != nil {
			invalidParams.AddNested("IncidentResponseTeam", err.(smithy.InvalidParamsError))
		}
	}
	if v.OptInFeatures != nil {
		if err := validateOptInFeatures(v.OptInFeatures); err != nil {
			invalidParams.AddNested("OptInFeatures", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCaseAttachmentDownloadUrlInput(v *GetCaseAttachmentDownloadUrlInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCaseAttachmentDownloadUrlInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if v.AttachmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttachmentId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCaseAttachmentUploadUrlInput(v *GetCaseAttachmentUploadUrlInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCaseAttachmentUploadUrlInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if v.FileName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileName"))
	}
	if v.ContentLength == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContentLength"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCaseInput(v *GetCaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCaseInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMembershipInput(v *GetMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMembershipInput"}
	if v.MembershipId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MembershipId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListCaseEditsInput(v *ListCaseEditsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListCaseEditsInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListCommentsInput(v *ListCommentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListCommentsInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCaseCommentInput(v *UpdateCaseCommentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCaseCommentInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if v.CommentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CommentId"))
	}
	if v.Body == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Body"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCaseInput(v *UpdateCaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCaseInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if v.WatchersToAdd != nil {
		if err := validateWatchers(v.WatchersToAdd); err != nil {
			invalidParams.AddNested("WatchersToAdd", err.(smithy.InvalidParamsError))
		}
	}
	if v.WatchersToDelete != nil {
		if err := validateWatchers(v.WatchersToDelete); err != nil {
			invalidParams.AddNested("WatchersToDelete", err.(smithy.InvalidParamsError))
		}
	}
	if v.ThreatActorIpAddressesToAdd != nil {
		if err := validateThreatActorIpList(v.ThreatActorIpAddressesToAdd); err != nil {
			invalidParams.AddNested("ThreatActorIpAddressesToAdd", err.(smithy.InvalidParamsError))
		}
	}
	if v.ThreatActorIpAddressesToDelete != nil {
		if err := validateThreatActorIpList(v.ThreatActorIpAddressesToDelete); err != nil {
			invalidParams.AddNested("ThreatActorIpAddressesToDelete", err.(smithy.InvalidParamsError))
		}
	}
	if v.ImpactedAwsRegionsToAdd != nil {
		if err := validateImpactedAwsRegionList(v.ImpactedAwsRegionsToAdd); err != nil {
			invalidParams.AddNested("ImpactedAwsRegionsToAdd", err.(smithy.InvalidParamsError))
		}
	}
	if v.ImpactedAwsRegionsToDelete != nil {
		if err := validateImpactedAwsRegionList(v.ImpactedAwsRegionsToDelete); err != nil {
			invalidParams.AddNested("ImpactedAwsRegionsToDelete", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCaseStatusInput(v *UpdateCaseStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCaseStatusInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if len(v.CaseStatus) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("CaseStatus"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateMembershipInput(v *UpdateMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMembershipInput"}
	if v.MembershipId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MembershipId"))
	}
	if v.IncidentResponseTeam != nil {
		if err := validateIncidentResponseTeam(v.IncidentResponseTeam); err != nil {
			invalidParams.AddNested("IncidentResponseTeam", err.(smithy.InvalidParamsError))
		}
	}
	if v.OptInFeatures != nil {
		if err := validateOptInFeatures(v.OptInFeatures); err != nil {
			invalidParams.AddNested("OptInFeatures", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateResolverTypeInput(v *UpdateResolverTypeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateResolverTypeInput"}
	if v.CaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CaseId"))
	}
	if len(v.ResolverType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
