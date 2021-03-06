// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateInstanceProfileInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance profile to create.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// InstanceProfileName is a required field
	InstanceProfileName *string `min:"1" type:"string" required:"true"`

	// The path to the instance profile. For more information about paths, see IAM
	// Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/).
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	Path *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateInstanceProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateInstanceProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateInstanceProfileInput"}

	if s.InstanceProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceProfileName"))
	}
	if s.InstanceProfileName != nil && len(*s.InstanceProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceProfileName", 1))
	}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful CreateInstanceProfile request.
type CreateInstanceProfileOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the new instance profile.
	//
	// InstanceProfile is a required field
	InstanceProfile *InstanceProfile `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateInstanceProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateInstanceProfile = "CreateInstanceProfile"

// CreateInstanceProfileRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates a new instance profile. For information about instance profiles,
// go to About Instance Profiles (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html).
//
// The number and size of IAM resources in an AWS account are limited. For more
// information, see IAM and STS Quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreateInstanceProfileRequest.
//    req := client.CreateInstanceProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateInstanceProfile
func (c *Client) CreateInstanceProfileRequest(input *CreateInstanceProfileInput) CreateInstanceProfileRequest {
	op := &aws.Operation{
		Name:       opCreateInstanceProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateInstanceProfileInput{}
	}

	req := c.newRequest(op, input, &CreateInstanceProfileOutput{})

	return CreateInstanceProfileRequest{Request: req, Input: input, Copy: c.CreateInstanceProfileRequest}
}

// CreateInstanceProfileRequest is the request type for the
// CreateInstanceProfile API operation.
type CreateInstanceProfileRequest struct {
	*aws.Request
	Input *CreateInstanceProfileInput
	Copy  func(*CreateInstanceProfileInput) CreateInstanceProfileRequest
}

// Send marshals and sends the CreateInstanceProfile API request.
func (r CreateInstanceProfileRequest) Send(ctx context.Context) (*CreateInstanceProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInstanceProfileResponse{
		CreateInstanceProfileOutput: r.Request.Data.(*CreateInstanceProfileOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInstanceProfileResponse is the response type for the
// CreateInstanceProfile API operation.
type CreateInstanceProfileResponse struct {
	*CreateInstanceProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInstanceProfile request.
func (r *CreateInstanceProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
