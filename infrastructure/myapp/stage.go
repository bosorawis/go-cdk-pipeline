package myapp

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

type AppStageProps struct {
	awscdk.StageProps
}

func AppStage(scope constructs.Construct, id string, props *AppStageProps) awscdk.Stage {
	var sprops awscdk.StageProps
	if props != nil {
		sprops = props.StageProps
	}
	stage := awscdk.NewStage(scope, &id, &sprops)
	NewApplicationStack(stage, "whatsmyip", nil)
	return stage
}
