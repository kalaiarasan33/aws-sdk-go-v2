// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetResourceDefinitionRequest
type GetResourceDefinitionInput struct {
	_ struct{} `type:"structure"`

	// ResourceDefinitionId is a required field
	ResourceDefinitionId *string `location:"uri" locationName:"ResourceDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetResourceDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourceDefinitionInput"}

	if s.ResourceDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourceDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ResourceDefinitionId != nil {
		v := *s.ResourceDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetResourceDefinitionResponse
type GetResourceDefinitionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	LastUpdatedTimestamp *string `type:"string"`

	LatestVersion *string `type:"string"`

	LatestVersionArn *string `type:"string"`

	Name *string `type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetResourceDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourceDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTimestamp != nil {
		v := *s.CreationTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedTimestamp != nil {
		v := *s.LastUpdatedTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastUpdatedTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersion != nil {
		v := *s.LatestVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LatestVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersionArn != nil {
		v := *s.LatestVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LatestVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetResourceDefinition = "GetResourceDefinition"

// GetResourceDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a resource definition, including its creation
// time and latest version.
//
//    // Example sending a request using GetResourceDefinitionRequest.
//    req := client.GetResourceDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetResourceDefinition
func (c *Client) GetResourceDefinitionRequest(input *GetResourceDefinitionInput) GetResourceDefinitionRequest {
	op := &aws.Operation{
		Name:       opGetResourceDefinition,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/resources/{ResourceDefinitionId}",
	}

	if input == nil {
		input = &GetResourceDefinitionInput{}
	}

	req := c.newRequest(op, input, &GetResourceDefinitionOutput{})
	return GetResourceDefinitionRequest{Request: req, Input: input, Copy: c.GetResourceDefinitionRequest}
}

// GetResourceDefinitionRequest is the request type for the
// GetResourceDefinition API operation.
type GetResourceDefinitionRequest struct {
	*aws.Request
	Input *GetResourceDefinitionInput
	Copy  func(*GetResourceDefinitionInput) GetResourceDefinitionRequest
}

// Send marshals and sends the GetResourceDefinition API request.
func (r GetResourceDefinitionRequest) Send(ctx context.Context) (*GetResourceDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourceDefinitionResponse{
		GetResourceDefinitionOutput: r.Request.Data.(*GetResourceDefinitionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetResourceDefinitionResponse is the response type for the
// GetResourceDefinition API operation.
type GetResourceDefinitionResponse struct {
	*GetResourceDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourceDefinition request.
func (r *GetResourceDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}