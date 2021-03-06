// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartRxNormInferenceJobInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the request. If you don't set the client request
	// token, Amazon Comprehend Medical generates one.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
	// (IAM) role that grants Amazon Comprehend Medical read access to your input
	// data. For more information, see Role-Based Permissions Required for Asynchronous
	// Operations (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions-med.html#auth-role-permissions-med).
	//
	// DataAccessRoleArn is a required field
	DataAccessRoleArn *string `min:"20" type:"string" required:"true"`

	// Specifies the format and location of the input data for the job.
	//
	// InputDataConfig is a required field
	InputDataConfig *InputDataConfig `type:"structure" required:"true"`

	// The identifier of the job.
	JobName *string `min:"1" type:"string"`

	// An AWS Key Management Service key to encrypt your output files. If you do
	// not specify a key, the files are written in plain text.
	KMSKey *string `min:"1" type:"string"`

	// The language of the input documents. All documents must be in the same language.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// Specifies where to send the output files.
	//
	// OutputDataConfig is a required field
	OutputDataConfig *OutputDataConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartRxNormInferenceJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartRxNormInferenceJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartRxNormInferenceJobInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DataAccessRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataAccessRoleArn"))
	}
	if s.DataAccessRoleArn != nil && len(*s.DataAccessRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("DataAccessRoleArn", 20))
	}

	if s.InputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDataConfig"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}
	if s.KMSKey != nil && len(*s.KMSKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KMSKey", 1))
	}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.OutputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputDataConfig"))
	}
	if s.InputDataConfig != nil {
		if err := s.InputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputDataConfig != nil {
		if err := s.OutputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartRxNormInferenceJobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the job.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartRxNormInferenceJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartRxNormInferenceJob = "StartRxNormInferenceJob"

// StartRxNormInferenceJobRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// Starts an asynchronous job to detect medication entities and link them to
// the RxNorm ontology. Use the DescribeRxNormInferenceJob operation to track
// the status of a job.
//
//    // Example sending a request using StartRxNormInferenceJobRequest.
//    req := client.StartRxNormInferenceJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/StartRxNormInferenceJob
func (c *Client) StartRxNormInferenceJobRequest(input *StartRxNormInferenceJobInput) StartRxNormInferenceJobRequest {
	op := &aws.Operation{
		Name:       opStartRxNormInferenceJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartRxNormInferenceJobInput{}
	}

	req := c.newRequest(op, input, &StartRxNormInferenceJobOutput{})

	return StartRxNormInferenceJobRequest{Request: req, Input: input, Copy: c.StartRxNormInferenceJobRequest}
}

// StartRxNormInferenceJobRequest is the request type for the
// StartRxNormInferenceJob API operation.
type StartRxNormInferenceJobRequest struct {
	*aws.Request
	Input *StartRxNormInferenceJobInput
	Copy  func(*StartRxNormInferenceJobInput) StartRxNormInferenceJobRequest
}

// Send marshals and sends the StartRxNormInferenceJob API request.
func (r StartRxNormInferenceJobRequest) Send(ctx context.Context) (*StartRxNormInferenceJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartRxNormInferenceJobResponse{
		StartRxNormInferenceJobOutput: r.Request.Data.(*StartRxNormInferenceJobOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartRxNormInferenceJobResponse is the response type for the
// StartRxNormInferenceJob API operation.
type StartRxNormInferenceJobResponse struct {
	*StartRxNormInferenceJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartRxNormInferenceJob request.
func (r *StartRxNormInferenceJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
