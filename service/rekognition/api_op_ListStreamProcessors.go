// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListStreamProcessorsInput struct {
	_ struct{} `type:"structure"`

	// Maximum number of stream processors you want Amazon Rekognition Video to
	// return in the response. The default is 1000.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous response was incomplete (because there are more stream processors
	// to retrieve), Amazon Rekognition Video returns a pagination token in the
	// response. You can use this pagination token to retrieve the next set of stream
	// processors.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListStreamProcessorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStreamProcessorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStreamProcessorsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListStreamProcessorsOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, Amazon Rekognition Video returns this token
	// that you can use in the subsequent request to retrieve the next set of stream
	// processors.
	NextToken *string `type:"string"`

	// List of stream processors that you have created.
	StreamProcessors []StreamProcessor `type:"list"`
}

// String returns the string representation
func (s ListStreamProcessorsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListStreamProcessors = "ListStreamProcessors"

// ListStreamProcessorsRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Gets a list of stream processors that you have created with CreateStreamProcessor.
//
//    // Example sending a request using ListStreamProcessorsRequest.
//    req := client.ListStreamProcessorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListStreamProcessorsRequest(input *ListStreamProcessorsInput) ListStreamProcessorsRequest {
	op := &aws.Operation{
		Name:       opListStreamProcessors,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListStreamProcessorsInput{}
	}

	req := c.newRequest(op, input, &ListStreamProcessorsOutput{})

	return ListStreamProcessorsRequest{Request: req, Input: input, Copy: c.ListStreamProcessorsRequest}
}

// ListStreamProcessorsRequest is the request type for the
// ListStreamProcessors API operation.
type ListStreamProcessorsRequest struct {
	*aws.Request
	Input *ListStreamProcessorsInput
	Copy  func(*ListStreamProcessorsInput) ListStreamProcessorsRequest
}

// Send marshals and sends the ListStreamProcessors API request.
func (r ListStreamProcessorsRequest) Send(ctx context.Context) (*ListStreamProcessorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStreamProcessorsResponse{
		ListStreamProcessorsOutput: r.Request.Data.(*ListStreamProcessorsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListStreamProcessorsRequestPaginator returns a paginator for ListStreamProcessors.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListStreamProcessorsRequest(input)
//   p := rekognition.NewListStreamProcessorsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListStreamProcessorsPaginator(req ListStreamProcessorsRequest) ListStreamProcessorsPaginator {
	return ListStreamProcessorsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListStreamProcessorsInput
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

// ListStreamProcessorsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListStreamProcessorsPaginator struct {
	aws.Pager
}

func (p *ListStreamProcessorsPaginator) CurrentPage() *ListStreamProcessorsOutput {
	return p.Pager.CurrentPage().(*ListStreamProcessorsOutput)
}

// ListStreamProcessorsResponse is the response type for the
// ListStreamProcessors API operation.
type ListStreamProcessorsResponse struct {
	*ListStreamProcessorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStreamProcessors request.
func (r *ListStreamProcessorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}