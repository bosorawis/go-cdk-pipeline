package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

type GoCdkPipelineStackProps struct {
	awscdk.StackProps
}

func main() {
	app := awscdk.NewApp(nil)

	NewPipelineStack(app, "GoPipelineStack", &GoCdkPipelineStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})
	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
