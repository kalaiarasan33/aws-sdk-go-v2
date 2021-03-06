// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Container for the parameters to the AcceptInboundCrossClusterSearchConnection
// operation.
type AcceptInboundCrossClusterSearchConnectionInput struct {
	_ struct{} `type:"structure"`

	// The id of the inbound connection that you want to accept.
	//
	// CrossClusterSearchConnectionId is a required field
	CrossClusterSearchConnectionId *string `location:"uri" locationName:"ConnectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptInboundCrossClusterSearchConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptInboundCrossClusterSearchConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptInboundCrossClusterSearchConnectionInput"}

	if s.CrossClusterSearchConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CrossClusterSearchConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptInboundCrossClusterSearchConnectionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CrossClusterSearchConnectionId != nil {
		v := *s.CrossClusterSearchConnectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConnectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a AcceptInboundCrossClusterSearchConnection operation. Contains
// details of accepted inbound connection.
type AcceptInboundCrossClusterSearchConnectionOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the InboundCrossClusterSearchConnection of accepted inbound connection.
	CrossClusterSearchConnection *InboundCrossClusterSearchConnection `type:"structure"`
}

// String returns the string representation
func (s AcceptInboundCrossClusterSearchConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptInboundCrossClusterSearchConnectionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CrossClusterSearchConnection != nil {
		v := s.CrossClusterSearchConnection

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CrossClusterSearchConnection", v, metadata)
	}
	return nil
}

const opAcceptInboundCrossClusterSearchConnection = "AcceptInboundCrossClusterSearchConnection"

// AcceptInboundCrossClusterSearchConnectionRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Allows the destination domain owner to accept an inbound cross-cluster search
// connection request.
//
//    // Example sending a request using AcceptInboundCrossClusterSearchConnectionRequest.
//    req := client.AcceptInboundCrossClusterSearchConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) AcceptInboundCrossClusterSearchConnectionRequest(input *AcceptInboundCrossClusterSearchConnectionInput) AcceptInboundCrossClusterSearchConnectionRequest {
	op := &aws.Operation{
		Name:       opAcceptInboundCrossClusterSearchConnection,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-01-01/es/ccs/inboundConnection/{ConnectionId}/accept",
	}

	if input == nil {
		input = &AcceptInboundCrossClusterSearchConnectionInput{}
	}

	req := c.newRequest(op, input, &AcceptInboundCrossClusterSearchConnectionOutput{})

	return AcceptInboundCrossClusterSearchConnectionRequest{Request: req, Input: input, Copy: c.AcceptInboundCrossClusterSearchConnectionRequest}
}

// AcceptInboundCrossClusterSearchConnectionRequest is the request type for the
// AcceptInboundCrossClusterSearchConnection API operation.
type AcceptInboundCrossClusterSearchConnectionRequest struct {
	*aws.Request
	Input *AcceptInboundCrossClusterSearchConnectionInput
	Copy  func(*AcceptInboundCrossClusterSearchConnectionInput) AcceptInboundCrossClusterSearchConnectionRequest
}

// Send marshals and sends the AcceptInboundCrossClusterSearchConnection API request.
func (r AcceptInboundCrossClusterSearchConnectionRequest) Send(ctx context.Context) (*AcceptInboundCrossClusterSearchConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptInboundCrossClusterSearchConnectionResponse{
		AcceptInboundCrossClusterSearchConnectionOutput: r.Request.Data.(*AcceptInboundCrossClusterSearchConnectionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptInboundCrossClusterSearchConnectionResponse is the response type for the
// AcceptInboundCrossClusterSearchConnection API operation.
type AcceptInboundCrossClusterSearchConnectionResponse struct {
	*AcceptInboundCrossClusterSearchConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptInboundCrossClusterSearchConnection request.
func (r *AcceptInboundCrossClusterSearchConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
