// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to create a new platform version.
type CreatePlatformVersionInput struct {
	_ struct{} `type:"structure"`

	// The name of the builder environment.
	EnvironmentName *string `min:"4" type:"string"`

	// The configuration option settings to apply to the builder environment.
	OptionSettings []ConfigurationOptionSetting `type:"list"`

	// The location of the platform definition archive in Amazon S3.
	//
	// PlatformDefinitionBundle is a required field
	PlatformDefinitionBundle *S3Location `type:"structure" required:"true"`

	// The name of your custom platform.
	//
	// PlatformName is a required field
	PlatformName *string `type:"string" required:"true"`

	// The number, such as 1.0.2, for the new platform version.
	//
	// PlatformVersion is a required field
	PlatformVersion *string `type:"string" required:"true"`

	// Specifies the tags applied to the new platform version.
	//
	// Elastic Beanstalk applies these tags only to the platform version. Environments
	// that you create using the platform version don't inherit the tags.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreatePlatformVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePlatformVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePlatformVersionInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}

	if s.PlatformDefinitionBundle == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlatformDefinitionBundle"))
	}

	if s.PlatformName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlatformName"))
	}

	if s.PlatformVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlatformVersion"))
	}
	if s.OptionSettings != nil {
		for i, v := range s.OptionSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OptionSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePlatformVersionOutput struct {
	_ struct{} `type:"structure"`

	// The builder used to create the custom platform.
	Builder *Builder `type:"structure"`

	// Detailed information about the new version of the custom platform.
	PlatformSummary *PlatformSummary `type:"structure"`
}

// String returns the string representation
func (s CreatePlatformVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePlatformVersion = "CreatePlatformVersion"

// CreatePlatformVersionRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Create a new version of your custom platform.
//
//    // Example sending a request using CreatePlatformVersionRequest.
//    req := client.CreatePlatformVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreatePlatformVersion
func (c *Client) CreatePlatformVersionRequest(input *CreatePlatformVersionInput) CreatePlatformVersionRequest {
	op := &aws.Operation{
		Name:       opCreatePlatformVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePlatformVersionInput{}
	}

	req := c.newRequest(op, input, &CreatePlatformVersionOutput{})

	return CreatePlatformVersionRequest{Request: req, Input: input, Copy: c.CreatePlatformVersionRequest}
}

// CreatePlatformVersionRequest is the request type for the
// CreatePlatformVersion API operation.
type CreatePlatformVersionRequest struct {
	*aws.Request
	Input *CreatePlatformVersionInput
	Copy  func(*CreatePlatformVersionInput) CreatePlatformVersionRequest
}

// Send marshals and sends the CreatePlatformVersion API request.
func (r CreatePlatformVersionRequest) Send(ctx context.Context) (*CreatePlatformVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePlatformVersionResponse{
		CreatePlatformVersionOutput: r.Request.Data.(*CreatePlatformVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePlatformVersionResponse is the response type for the
// CreatePlatformVersion API operation.
type CreatePlatformVersionResponse struct {
	*CreatePlatformVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePlatformVersion request.
func (r *CreatePlatformVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
