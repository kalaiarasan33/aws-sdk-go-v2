// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The example tests how requests and responses are serialized when there's no
// request or response payload because the operation has no input or output. While
// this should be rare, code generators must support this.
func (c *Client) NoInputAndNoOutput(ctx context.Context, params *NoInputAndNoOutputInput, optFns ...func(*Options)) (*NoInputAndNoOutputOutput, error) {
	stack := middleware.NewStack("NoInputAndNoOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	stack.Initialize.Add(newServiceMetadataMiddleware_opNoInputAndNoOutput(options.Region), middleware.Before)
	addawsRestjson1_serdeOpNoInputAndNoOutputMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "NoInputAndNoOutput",
			Err:           err,
		}
	}
	out := result.(*NoInputAndNoOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NoInputAndNoOutputInput struct {
}

type NoInputAndNoOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpNoInputAndNoOutputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpNoInputAndNoOutput{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpNoInputAndNoOutput{}, middleware.After)
}

func newServiceMetadataMiddleware_opNoInputAndNoOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "NoInputAndNoOutput",
	}
}
