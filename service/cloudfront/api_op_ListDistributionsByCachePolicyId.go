// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDistributionsByCachePolicyIdInput struct {
	_ struct{} `type:"structure"`

	// The ID of the cache policy whose associated distribution IDs you want to
	// list.
	//
	// CachePolicyId is a required field
	CachePolicyId *string `location:"uri" locationName:"CachePolicyId" type:"string" required:"true"`

	// Use this field when paginating results to indicate where to begin in your
	// list of distribution IDs. The response includes distribution IDs in the list
	// that occur after the marker. To get the next page of the list, set this field’s
	// value to the value of NextMarker from the current page’s response.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of distribution IDs that you want in the response.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`
}

// String returns the string representation
func (s ListDistributionsByCachePolicyIdInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDistributionsByCachePolicyIdInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDistributionsByCachePolicyIdInput"}

	if s.CachePolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CachePolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionsByCachePolicyIdInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CachePolicyId != nil {
		v := *s.CachePolicyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "CachePolicyId", protocol.StringValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

type ListDistributionsByCachePolicyIdOutput struct {
	_ struct{} `type:"structure" payload:"DistributionIdList"`

	// A list of distribution IDs.
	DistributionIdList *DistributionIdList `type:"structure"`
}

// String returns the string representation
func (s ListDistributionsByCachePolicyIdOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionsByCachePolicyIdOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DistributionIdList != nil {
		v := s.DistributionIdList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "DistributionIdList", v, metadata)
	}
	return nil
}

const opListDistributionsByCachePolicyId = "ListDistributionsByCachePolicyId2020_05_31"

// ListDistributionsByCachePolicyIdRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Gets a list of distribution IDs for distributions that have a cache behavior
// that’s associated with the specified cache policy.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that
// you specify, or the default maximum, the response is paginated. To get the
// next page of items, send a subsequent request that specifies the NextMarker
// value from the current response as the Marker value in the subsequent request.
//
//    // Example sending a request using ListDistributionsByCachePolicyIdRequest.
//    req := client.ListDistributionsByCachePolicyIdRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/ListDistributionsByCachePolicyId
func (c *Client) ListDistributionsByCachePolicyIdRequest(input *ListDistributionsByCachePolicyIdInput) ListDistributionsByCachePolicyIdRequest {
	op := &aws.Operation{
		Name:       opListDistributionsByCachePolicyId,
		HTTPMethod: "GET",
		HTTPPath:   "/2020-05-31/distributionsByCachePolicyId/{CachePolicyId}",
	}

	if input == nil {
		input = &ListDistributionsByCachePolicyIdInput{}
	}

	req := c.newRequest(op, input, &ListDistributionsByCachePolicyIdOutput{})

	return ListDistributionsByCachePolicyIdRequest{Request: req, Input: input, Copy: c.ListDistributionsByCachePolicyIdRequest}
}

// ListDistributionsByCachePolicyIdRequest is the request type for the
// ListDistributionsByCachePolicyId API operation.
type ListDistributionsByCachePolicyIdRequest struct {
	*aws.Request
	Input *ListDistributionsByCachePolicyIdInput
	Copy  func(*ListDistributionsByCachePolicyIdInput) ListDistributionsByCachePolicyIdRequest
}

// Send marshals and sends the ListDistributionsByCachePolicyId API request.
func (r ListDistributionsByCachePolicyIdRequest) Send(ctx context.Context) (*ListDistributionsByCachePolicyIdResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDistributionsByCachePolicyIdResponse{
		ListDistributionsByCachePolicyIdOutput: r.Request.Data.(*ListDistributionsByCachePolicyIdOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDistributionsByCachePolicyIdResponse is the response type for the
// ListDistributionsByCachePolicyId API operation.
type ListDistributionsByCachePolicyIdResponse struct {
	*ListDistributionsByCachePolicyIdOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDistributionsByCachePolicyId request.
func (r *ListDistributionsByCachePolicyIdResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
