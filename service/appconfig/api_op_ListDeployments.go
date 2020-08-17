// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDeploymentsInput struct {
	_ struct{} `type:"structure"`

	// The application ID.
	//
	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"ApplicationId" type:"string" required:"true"`

	// The environment ID.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `location:"uri" locationName:"EnvironmentId" type:"string" required:"true"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `location:"querystring" locationName:"max_results" min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `location:"querystring" locationName:"next_token" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDeploymentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeploymentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeploymentsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.EnvironmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeploymentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnvironmentId != nil {
		v := *s.EnvironmentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EnvironmentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max_results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next_token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListDeploymentsOutput struct {
	_ struct{} `type:"structure"`

	// The elements from this collection.
	Items []DeploymentSummary `type:"list"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDeploymentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeploymentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDeployments = "ListDeployments"

// ListDeploymentsRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// Lists the deployments for an environment.
//
//    // Example sending a request using ListDeploymentsRequest.
//    req := client.ListDeploymentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/ListDeployments
func (c *Client) ListDeploymentsRequest(input *ListDeploymentsInput) ListDeploymentsRequest {
	op := &aws.Operation{
		Name:       opListDeployments,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{ApplicationId}/environments/{EnvironmentId}/deployments",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDeploymentsInput{}
	}

	req := c.newRequest(op, input, &ListDeploymentsOutput{})

	return ListDeploymentsRequest{Request: req, Input: input, Copy: c.ListDeploymentsRequest}
}

// ListDeploymentsRequest is the request type for the
// ListDeployments API operation.
type ListDeploymentsRequest struct {
	*aws.Request
	Input *ListDeploymentsInput
	Copy  func(*ListDeploymentsInput) ListDeploymentsRequest
}

// Send marshals and sends the ListDeployments API request.
func (r ListDeploymentsRequest) Send(ctx context.Context) (*ListDeploymentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeploymentsResponse{
		ListDeploymentsOutput: r.Request.Data.(*ListDeploymentsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDeploymentsRequestPaginator returns a paginator for ListDeployments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDeploymentsRequest(input)
//   p := appconfig.NewListDeploymentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDeploymentsPaginator(req ListDeploymentsRequest) ListDeploymentsPaginator {
	return ListDeploymentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDeploymentsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDeploymentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDeploymentsPaginator struct {
	aws.Pager
}

func (p *ListDeploymentsPaginator) CurrentPage() *ListDeploymentsOutput {
	return p.Pager.CurrentPage().(*ListDeploymentsOutput)
}

// ListDeploymentsResponse is the response type for the
// ListDeployments API operation.
type ListDeploymentsResponse struct {
	*ListDeploymentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeployments request.
func (r *ListDeploymentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}