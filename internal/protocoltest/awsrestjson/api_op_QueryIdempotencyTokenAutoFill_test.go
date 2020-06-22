// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestClient_QueryIdempotencyTokenAutoFill_awsRestjson1Serialize(t *testing.T) {
	origRandReader := randReader
	defer func() { randReader = origRandReader }()
	randReader = &smithytesting.ByteLoop{}

	cases := map[string]struct {
		Params        *QueryIdempotencyTokenAutoFillInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Automatically adds idempotency token when not set
		"RestJsonQueryIdempotencyTokenAutoFill": {
			Params:        &QueryIdempotencyTokenAutoFillInput{},
			ExpectMethod:  "POST",
			ExpectURIPath: "/QueryIdempotencyTokenAutoFill",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "token", Value: "00000000-0000-4000-8000-000000000000"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Uses the given idempotency token as-is
		"RestJsonQueryIdempotencyTokenAutoFillIsSet": {
			Params: &QueryIdempotencyTokenAutoFillInput{
				Token: ptr.String("00000000-0000-4000-8000-000000000000"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/QueryIdempotencyTokenAutoFill",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "token", Value: "00000000-0000-4000-8000-000000000000"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					actualReq = r
					return &http.Response{
						StatusCode: 200,
						Header:     http.Header{},
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}, nil
				}),
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = "https://127.0.0.1"
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2"})
			result, err := client.QueryIdempotencyTokenAutoFill(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.Path; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if actualReq.Body != nil {
				defer actualReq.Body.Close()
			}
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}
