// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Gets the value that Amazon Route 53 returns in response to a DNS request
// for a specified record name and type. You can optionally specify the IP address
// of a DNS resolver, an EDNS0 client subnet IP address, and a subnet mask.
type TestDNSAnswerInput struct {
	_ struct{} `type:"structure"`

	// If the resolver that you specified for resolverip supports EDNS0, specify
	// the IPv4 or IPv6 address of a client in the applicable location, for example,
	// 192.0.2.44 or 2001:db8:85a3::8a2e:370:7334.
	EDNS0ClientSubnetIP *string `location:"querystring" locationName:"edns0clientsubnetip" type:"string"`

	// If you specify an IP address for edns0clientsubnetip, you can optionally
	// specify the number of bits of the IP address that you want the checking tool
	// to include in the DNS query. For example, if you specify 192.0.2.44 for edns0clientsubnetip
	// and 24 for edns0clientsubnetmask, the checking tool will simulate a request
	// from 192.0.2.0/24. The default value is 24 bits for IPv4 addresses and 64
	// bits for IPv6 addresses.
	//
	// The range of valid values depends on whether edns0clientsubnetip is an IPv4
	// or an IPv6 address:
	//
	//    * IPv4: Specify a value between 0 and 32
	//
	//    * IPv6: Specify a value between 0 and 128
	EDNS0ClientSubnetMask *string `location:"querystring" locationName:"edns0clientsubnetmask" type:"string"`

	// The ID of the hosted zone that you want Amazon Route 53 to simulate a query
	// for.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `location:"querystring" locationName:"hostedzoneid" type:"string" required:"true"`

	// The name of the resource record set that you want Amazon Route 53 to simulate
	// a query for.
	//
	// RecordName is a required field
	RecordName *string `location:"querystring" locationName:"recordname" type:"string" required:"true"`

	// The type of the resource record set.
	//
	// RecordType is a required field
	RecordType RRType `location:"querystring" locationName:"recordtype" type:"string" required:"true" enum:"true"`

	// If you want to simulate a request from a specific DNS resolver, specify the
	// IP address for that resolver. If you omit this value, TestDnsAnswer uses
	// the IP address of a DNS resolver in the AWS US East (N. Virginia) Region
	// (us-east-1).
	ResolverIP *string `location:"querystring" locationName:"resolverip" type:"string"`
}

// String returns the string representation
func (s TestDNSAnswerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestDNSAnswerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestDNSAnswerInput"}

	if s.HostedZoneId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostedZoneId"))
	}

	if s.RecordName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RecordName"))
	}
	if len(s.RecordType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RecordType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TestDNSAnswerInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.EDNS0ClientSubnetIP != nil {
		v := *s.EDNS0ClientSubnetIP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "edns0clientsubnetip", protocol.StringValue(v), metadata)
	}
	if s.EDNS0ClientSubnetMask != nil {
		v := *s.EDNS0ClientSubnetMask

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "edns0clientsubnetmask", protocol.StringValue(v), metadata)
	}
	if s.HostedZoneId != nil {
		v := *s.HostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "hostedzoneid", protocol.StringValue(v), metadata)
	}
	if s.RecordName != nil {
		v := *s.RecordName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "recordname", protocol.StringValue(v), metadata)
	}
	if len(s.RecordType) > 0 {
		v := s.RecordType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "recordtype", v, metadata)
	}
	if s.ResolverIP != nil {
		v := *s.ResolverIP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resolverip", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response to a TestDNSAnswer request.
type TestDNSAnswerOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Route 53 name server used to respond to the request.
	//
	// Nameserver is a required field
	Nameserver *string `type:"string" required:"true"`

	// The protocol that Amazon Route 53 used to respond to the request, either
	// UDP or TCP.
	//
	// Protocol is a required field
	Protocol *string `type:"string" required:"true"`

	// A list that contains values that Amazon Route 53 returned for this resource
	// record set.
	//
	// RecordData is a required field
	RecordData []string `locationNameList:"RecordDataEntry" type:"list" required:"true"`

	// The name of the resource record set that you submitted a request for.
	//
	// RecordName is a required field
	RecordName *string `type:"string" required:"true"`

	// The type of the resource record set that you submitted a request for.
	//
	// RecordType is a required field
	RecordType RRType `type:"string" required:"true" enum:"true"`

	// A code that indicates whether the request is valid or not. The most common
	// response code is NOERROR, meaning that the request is valid. If the response
	// is not valid, Amazon Route 53 returns a response code that describes the
	// error. For a list of possible response codes, see DNS RCODES (http://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6)
	// on the IANA website.
	//
	// ResponseCode is a required field
	ResponseCode *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TestDNSAnswerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TestDNSAnswerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Nameserver != nil {
		v := *s.Nameserver

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Nameserver", protocol.StringValue(v), metadata)
	}
	if s.Protocol != nil {
		v := *s.Protocol

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Protocol", protocol.StringValue(v), metadata)
	}
	if s.RecordData != nil {
		v := s.RecordData

		metadata := protocol.Metadata{ListLocationName: "RecordDataEntry"}
		ls0 := e.List(protocol.BodyTarget, "RecordData", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.StringValue(v1))
		}
		ls0.End()

	}
	if s.RecordName != nil {
		v := *s.RecordName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RecordName", protocol.StringValue(v), metadata)
	}
	if len(s.RecordType) > 0 {
		v := s.RecordType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RecordType", v, metadata)
	}
	if s.ResponseCode != nil {
		v := *s.ResponseCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResponseCode", protocol.StringValue(v), metadata)
	}
	return nil
}

const opTestDNSAnswer = "TestDNSAnswer"

// TestDNSAnswerRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets the value that Amazon Route 53 returns in response to a DNS request
// for a specified record name and type. You can optionally specify the IP address
// of a DNS resolver, an EDNS0 client subnet IP address, and a subnet mask.
//
//    // Example sending a request using TestDNSAnswerRequest.
//    req := client.TestDNSAnswerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/TestDNSAnswer
func (c *Client) TestDNSAnswerRequest(input *TestDNSAnswerInput) TestDNSAnswerRequest {
	op := &aws.Operation{
		Name:       opTestDNSAnswer,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/testdnsanswer",
	}

	if input == nil {
		input = &TestDNSAnswerInput{}
	}

	req := c.newRequest(op, input, &TestDNSAnswerOutput{})

	return TestDNSAnswerRequest{Request: req, Input: input, Copy: c.TestDNSAnswerRequest}
}

// TestDNSAnswerRequest is the request type for the
// TestDNSAnswer API operation.
type TestDNSAnswerRequest struct {
	*aws.Request
	Input *TestDNSAnswerInput
	Copy  func(*TestDNSAnswerInput) TestDNSAnswerRequest
}

// Send marshals and sends the TestDNSAnswer API request.
func (r TestDNSAnswerRequest) Send(ctx context.Context) (*TestDNSAnswerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestDNSAnswerResponse{
		TestDNSAnswerOutput: r.Request.Data.(*TestDNSAnswerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestDNSAnswerResponse is the response type for the
// TestDNSAnswer API operation.
type TestDNSAnswerResponse struct {
	*TestDNSAnswerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestDNSAnswer request.
func (r *TestDNSAnswerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}