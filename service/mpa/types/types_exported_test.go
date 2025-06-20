// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/mpa/types"
)

func ExampleApprovalStrategy_outputUsage() {
	var union types.ApprovalStrategy
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ApprovalStrategyMemberMofN:
		_ = v.Value // Value is types.MofNApprovalStrategy

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MofNApprovalStrategy

func ExampleApprovalStrategyResponse_outputUsage() {
	var union types.ApprovalStrategyResponse
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ApprovalStrategyResponseMemberMofN:
		_ = v.Value // Value is types.MofNApprovalStrategy

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MofNApprovalStrategy

func ExampleIdentitySourceParametersForGet_outputUsage() {
	var union types.IdentitySourceParametersForGet
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.IdentitySourceParametersForGetMemberIamIdentityCenter:
		_ = v.Value // Value is types.IamIdentityCenterForGet

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.IamIdentityCenterForGet

func ExampleIdentitySourceParametersForList_outputUsage() {
	var union types.IdentitySourceParametersForList
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.IdentitySourceParametersForListMemberIamIdentityCenter:
		_ = v.Value // Value is types.IamIdentityCenterForList

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.IamIdentityCenterForList
