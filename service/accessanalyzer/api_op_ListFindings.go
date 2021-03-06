// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Retrieves a list of findings generated by the specified analyzer.
type ListFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the analyzer to retrieve findings from.
	//
	// AnalyzerArn is a required field
	AnalyzerArn *string `locationName:"analyzerArn" type:"string" required:"true"`

	// A filter to match for the findings to return.
	Filter map[string]Criterion `locationName:"filter" type:"map"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// A token used for pagination of results returned.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The sort order for the findings returned.
	Sort *SortCriteria `locationName:"sort" type:"structure"`
}

// String returns the string representation
func (s ListFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFindingsInput"}

	if s.AnalyzerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AnalyzerArn"))
	}
	if s.Filter != nil {
		for i, v := range s.Filter {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filter", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AnalyzerArn != nil {
		v := *s.AnalyzerArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "analyzerArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Filter != nil {
		v := s.Filter

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "filter", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Sort != nil {
		v := s.Sort

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "sort", v, metadata)
	}
	return nil
}

// The response to the request.
type ListFindingsOutput struct {
	_ struct{} `type:"structure"`

	// A list of findings retrieved from the analyzer that match the filter criteria
	// specified, if any.
	//
	// Findings is a required field
	Findings []FindingSummary `locationName:"findings" type:"list" required:"true"`

	// A token used for pagination of results returned.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Findings != nil {
		v := s.Findings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFindings = "ListFindings"

// ListFindingsRequest returns a request value for making API operation for
// Access Analyzer.
//
// Retrieves a list of findings generated by the specified analyzer.
//
//    // Example sending a request using ListFindingsRequest.
//    req := client.ListFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/accessanalyzer-2019-11-01/ListFindings
func (c *Client) ListFindingsRequest(input *ListFindingsInput) ListFindingsRequest {
	op := &aws.Operation{
		Name:       opListFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/finding",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFindingsInput{}
	}

	req := c.newRequest(op, input, &ListFindingsOutput{})

	return ListFindingsRequest{Request: req, Input: input, Copy: c.ListFindingsRequest}
}

// ListFindingsRequest is the request type for the
// ListFindings API operation.
type ListFindingsRequest struct {
	*aws.Request
	Input *ListFindingsInput
	Copy  func(*ListFindingsInput) ListFindingsRequest
}

// Send marshals and sends the ListFindings API request.
func (r ListFindingsRequest) Send(ctx context.Context) (*ListFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFindingsResponse{
		ListFindingsOutput: r.Request.Data.(*ListFindingsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFindingsRequestPaginator returns a paginator for ListFindings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFindingsRequest(input)
//   p := accessanalyzer.NewListFindingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFindingsPaginator(req ListFindingsRequest) ListFindingsPaginator {
	return ListFindingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFindingsInput
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

// ListFindingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFindingsPaginator struct {
	aws.Pager
}

func (p *ListFindingsPaginator) CurrentPage() *ListFindingsOutput {
	return p.Pager.CurrentPage().(*ListFindingsOutput)
}

// ListFindingsResponse is the response type for the
// ListFindings API operation.
type ListFindingsResponse struct {
	*ListFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFindings request.
func (r *ListFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
