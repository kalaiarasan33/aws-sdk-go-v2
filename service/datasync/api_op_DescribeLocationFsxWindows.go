// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLocationFsxWindowsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the FSx for Windows location to describe.
	//
	// LocationArn is a required field
	LocationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLocationFsxWindowsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLocationFsxWindowsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLocationFsxWindowsInput"}

	if s.LocationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LocationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeLocationFsxWindowsOutput struct {
	_ struct{} `type:"structure"`

	// The time that the FSx for Windows location was created.
	CreationTime *time.Time `type:"timestamp"`

	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain *string `type:"string"`

	// The Amazon resource Name (ARN) of the FSx for Windows location that was described.
	LocationArn *string `type:"string"`

	// The URL of the FSx for Windows location that was described.
	LocationUri *string `type:"string"`

	// The Amazon Resource Names (ARNs) of the security groups that are configured
	// for the for the FSx for Windows file system.
	SecurityGroupArns []string `min:"1" type:"list"`

	// The user who has the permissions to access files and folders in the FSx for
	// Windows file system.
	User *string `type:"string"`
}

// String returns the string representation
func (s DescribeLocationFsxWindowsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLocationFsxWindows = "DescribeLocationFsxWindows"

// DescribeLocationFsxWindowsRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns metadata, such as the path information about an Amazon FSx for Windows
// location.
//
//    // Example sending a request using DescribeLocationFsxWindowsRequest.
//    req := client.DescribeLocationFsxWindowsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationFsxWindows
func (c *Client) DescribeLocationFsxWindowsRequest(input *DescribeLocationFsxWindowsInput) DescribeLocationFsxWindowsRequest {
	op := &aws.Operation{
		Name:       opDescribeLocationFsxWindows,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLocationFsxWindowsInput{}
	}

	req := c.newRequest(op, input, &DescribeLocationFsxWindowsOutput{})

	return DescribeLocationFsxWindowsRequest{Request: req, Input: input, Copy: c.DescribeLocationFsxWindowsRequest}
}

// DescribeLocationFsxWindowsRequest is the request type for the
// DescribeLocationFsxWindows API operation.
type DescribeLocationFsxWindowsRequest struct {
	*aws.Request
	Input *DescribeLocationFsxWindowsInput
	Copy  func(*DescribeLocationFsxWindowsInput) DescribeLocationFsxWindowsRequest
}

// Send marshals and sends the DescribeLocationFsxWindows API request.
func (r DescribeLocationFsxWindowsRequest) Send(ctx context.Context) (*DescribeLocationFsxWindowsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLocationFsxWindowsResponse{
		DescribeLocationFsxWindowsOutput: r.Request.Data.(*DescribeLocationFsxWindowsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLocationFsxWindowsResponse is the response type for the
// DescribeLocationFsxWindows API operation.
type DescribeLocationFsxWindowsResponse struct {
	*DescribeLocationFsxWindowsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLocationFsxWindows request.
func (r *DescribeLocationFsxWindowsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
