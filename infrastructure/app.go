package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)


type AppStackProps struct {
	awscdk.StackProps
}

func NewApplicationStack(scope constructs.Construct, id string, props *AppStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	lambda := awslambda.NewFunction(stack, jsii.String("myFunction"), &awslambda.FunctionProps{
		Handler: jsii.String("handler"),
		Runtime: awslambda.Runtime_GO_1_X(),
		Timeout: awscdk.Duration_Seconds(jsii.Number(3.0)),
		Code: awslambda.Code_FromAsset(jsii.String("./bin/lambda"), nil),
	})

	api := awsapigateway.NewRestApi(stack, jsii.String("my-api"), &awsapigateway.RestApiProps{
		DeployOptions: &awsapigateway.StageOptions{
			StageName: jsii.String("v1"),
			MetricsEnabled: jsii.Bool(true),
			LoggingLevel: awsapigateway.MethodLoggingLevel_INFO,
			DataTraceEnabled: jsii.Bool(true),
		},
	})
	api.Root().AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(lambda, nil),
		nil,
	)

	awscdk.NewCfnOutput(stack, jsii.String("ApiUrl"), &awscdk.CfnOutputProps{
		Value: api.Url(),
	})
	return stack
}
