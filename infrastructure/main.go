package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

func main() {
	app := awscdk.NewApp(nil)

	NewPipelineStack(app, "GoPipelineStack", &PipelineStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})
	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
