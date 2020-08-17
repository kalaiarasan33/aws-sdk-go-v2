// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RequestServiceQuotaIncreaseInput struct {
	_ struct{} `type:"structure"`

	// Specifies the value submitted in the service quota increase request.
	//
	// DesiredValue is a required field
	DesiredValue *float64 `type:"double" required:"true"`

	// Specifies the service quota that you want to use.
	//
	// QuotaCode is a required field
	QuotaCode *string `min:"1" type:"string" required:"true"`

	// Specifies the service that you want to use.
	//
	// ServiceCode is a required field
	ServiceCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RequestServiceQuotaIncreaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestServiceQuotaIncreaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestServiceQuotaIncreaseInput"}

	if s.DesiredValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("DesiredValue"))
	}

	if s.QuotaCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("QuotaCode"))
	}
	if s.QuotaCode != nil && len(*s.QuotaCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QuotaCode", 1))
	}

	if s.ServiceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceCode"))
	}
	if s.ServiceCode != nil && len(*s.ServiceCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RequestServiceQuotaIncreaseOutput struct {
	_ struct{} `type:"structure"`

	// Returns a list of service quota requests.
	RequestedQuota *RequestedServiceQuotaChange `type:"structure"`
}

// String returns the string representation
func (s RequestServiceQuotaIncreaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opRequestServiceQuotaIncrease = "RequestServiceQuotaIncrease"

// RequestServiceQuotaIncreaseRequest returns a request value for making API operation for
// Service Quotas.
//
// Retrieves the details of a service quota increase request. The response to
// this command provides the details in the RequestedServiceQuotaChange object.
//
//    // Example sending a request using RequestServiceQuotaIncreaseRequest.
//    req := client.RequestServiceQuotaIncreaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/RequestServiceQuotaIncrease
func (c *Client) RequestServiceQuotaIncreaseRequest(input *RequestServiceQuotaIncreaseInput) RequestServiceQuotaIncreaseRequest {
	op := &aws.Operation{
		Name:       opRequestServiceQuotaIncrease,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RequestServiceQuotaIncreaseInput{}
	}

	req := c.newRequest(op, input, &RequestServiceQuotaIncreaseOutput{})

	return RequestServiceQuotaIncreaseRequest{Request: req, Input: input, Copy: c.RequestServiceQuotaIncreaseRequest}
}

// RequestServiceQuotaIncreaseRequest is the request type for the
// RequestServiceQuotaIncrease API operation.
type RequestServiceQuotaIncreaseRequest struct {
	*aws.Request
	Input *RequestServiceQuotaIncreaseInput
	Copy  func(*RequestServiceQuotaIncreaseInput) RequestServiceQuotaIncreaseRequest
}

// Send marshals and sends the RequestServiceQuotaIncrease API request.
func (r RequestServiceQuotaIncreaseRequest) Send(ctx context.Context) (*RequestServiceQuotaIncreaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RequestServiceQuotaIncreaseResponse{
		RequestServiceQuotaIncreaseOutput: r.Request.Data.(*RequestServiceQuotaIncreaseOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RequestServiceQuotaIncreaseResponse is the response type for the
// RequestServiceQuotaIncrease API operation.
type RequestServiceQuotaIncreaseResponse struct {
	*RequestServiceQuotaIncreaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RequestServiceQuotaIncrease request.
func (r *RequestServiceQuotaIncreaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}