// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SearchNetworkProfilesInput struct {
	_ struct{} `type:"structure"`

	// The filters to use to list a specified set of network profiles. Valid filters
	// are NetworkProfileName, Ssid, and SecurityType.
	Filters []Filter `type:"list"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `min:"1" type:"string"`

	// The sort order to use to list the specified set of network profiles. Valid
	// sort criteria includes NetworkProfileName, Ssid, and SecurityType.
	SortCriteria []Sort `type:"list"`
}

// String returns the string representation
func (s SearchNetworkProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchNetworkProfilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchNetworkProfilesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SortCriteria != nil {
		for i, v := range s.SortCriteria {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SortCriteria", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SearchNetworkProfilesOutput struct {
	_ struct{} `type:"structure"`

	// The network profiles that meet the specified set of filter criteria, in sort
	// order. It is a list of NetworkProfileData objects.
	NetworkProfiles []NetworkProfileData `type:"list"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `min:"1" type:"string"`

	// The total number of network profiles returned.
	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s SearchNetworkProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchNetworkProfiles = "SearchNetworkProfiles"

// SearchNetworkProfilesRequest returns a request value for making API operation for
// Alexa For Business.
//
// Searches network profiles and lists the ones that meet a set of filter and
// sort criteria.
//
//    // Example sending a request using SearchNetworkProfilesRequest.
//    req := client.SearchNetworkProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SearchNetworkProfiles
func (c *Client) SearchNetworkProfilesRequest(input *SearchNetworkProfilesInput) SearchNetworkProfilesRequest {
	op := &aws.Operation{
		Name:       opSearchNetworkProfiles,
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
		input = &SearchNetworkProfilesInput{}
	}

	req := c.newRequest(op, input, &SearchNetworkProfilesOutput{})

	return SearchNetworkProfilesRequest{Request: req, Input: input, Copy: c.SearchNetworkProfilesRequest}
}

// SearchNetworkProfilesRequest is the request type for the
// SearchNetworkProfiles API operation.
type SearchNetworkProfilesRequest struct {
	*aws.Request
	Input *SearchNetworkProfilesInput
	Copy  func(*SearchNetworkProfilesInput) SearchNetworkProfilesRequest
}

// Send marshals and sends the SearchNetworkProfiles API request.
func (r SearchNetworkProfilesRequest) Send(ctx context.Context) (*SearchNetworkProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchNetworkProfilesResponse{
		SearchNetworkProfilesOutput: r.Request.Data.(*SearchNetworkProfilesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchNetworkProfilesRequestPaginator returns a paginator for SearchNetworkProfiles.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchNetworkProfilesRequest(input)
//   p := alexaforbusiness.NewSearchNetworkProfilesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchNetworkProfilesPaginator(req SearchNetworkProfilesRequest) SearchNetworkProfilesPaginator {
	return SearchNetworkProfilesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchNetworkProfilesInput
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

// SearchNetworkProfilesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchNetworkProfilesPaginator struct {
	aws.Pager
}

func (p *SearchNetworkProfilesPaginator) CurrentPage() *SearchNetworkProfilesOutput {
	return p.Pager.CurrentPage().(*SearchNetworkProfilesOutput)
}

// SearchNetworkProfilesResponse is the response type for the
// SearchNetworkProfiles API operation.
type SearchNetworkProfilesResponse struct {
	*SearchNetworkProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchNetworkProfiles request.
func (r *SearchNetworkProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}