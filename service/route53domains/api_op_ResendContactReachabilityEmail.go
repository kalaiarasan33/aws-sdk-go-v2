// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResendContactReachabilityEmailInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain for which you want Route 53 to resend a confirmation
	// email to the registrant contact.
	DomainName *string `locationName:"domainName" type:"string"`
}

// String returns the string representation
func (s ResendContactReachabilityEmailInput) String() string {
	return awsutil.Prettify(s)
}

type ResendContactReachabilityEmailOutput struct {
	_ struct{} `type:"structure"`

	// The domain name for which you requested a confirmation email.
	DomainName *string `locationName:"domainName" type:"string"`

	// The email address for the registrant contact at the time that we sent the
	// verification email.
	EmailAddress *string `locationName:"emailAddress" type:"string"`

	// True if the email address for the registrant contact has already been verified,
	// and false otherwise. If the email address has already been verified, we don't
	// send another confirmation email.
	IsAlreadyVerified *bool `locationName:"isAlreadyVerified" type:"boolean"`
}

// String returns the string representation
func (s ResendContactReachabilityEmailOutput) String() string {
	return awsutil.Prettify(s)
}

const opResendContactReachabilityEmail = "ResendContactReachabilityEmail"

// ResendContactReachabilityEmailRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// For operations that require confirmation that the email address for the registrant
// contact is valid, such as registering a new domain, this operation resends
// the confirmation email to the current email address for the registrant contact.
//
//    // Example sending a request using ResendContactReachabilityEmailRequest.
//    req := client.ResendContactReachabilityEmailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/ResendContactReachabilityEmail
func (c *Client) ResendContactReachabilityEmailRequest(input *ResendContactReachabilityEmailInput) ResendContactReachabilityEmailRequest {
	op := &aws.Operation{
		Name:       opResendContactReachabilityEmail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResendContactReachabilityEmailInput{}
	}

	req := c.newRequest(op, input, &ResendContactReachabilityEmailOutput{})

	return ResendContactReachabilityEmailRequest{Request: req, Input: input, Copy: c.ResendContactReachabilityEmailRequest}
}

// ResendContactReachabilityEmailRequest is the request type for the
// ResendContactReachabilityEmail API operation.
type ResendContactReachabilityEmailRequest struct {
	*aws.Request
	Input *ResendContactReachabilityEmailInput
	Copy  func(*ResendContactReachabilityEmailInput) ResendContactReachabilityEmailRequest
}

// Send marshals and sends the ResendContactReachabilityEmail API request.
func (r ResendContactReachabilityEmailRequest) Send(ctx context.Context) (*ResendContactReachabilityEmailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResendContactReachabilityEmailResponse{
		ResendContactReachabilityEmailOutput: r.Request.Data.(*ResendContactReachabilityEmailOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResendContactReachabilityEmailResponse is the response type for the
// ResendContactReachabilityEmail API operation.
type ResendContactReachabilityEmailResponse struct {
	*ResendContactReachabilityEmailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResendContactReachabilityEmail request.
func (r *ResendContactReachabilityEmailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}