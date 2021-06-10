package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipelineactions"
	"github.com/aws/aws-cdk-go/awscdk/awsssm"
	"github.com/aws/aws-cdk-go/awscdk/pipelines"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"go-cdk-pipeline/infrastructure/myapp"
)

type PipelineStackProps struct {
	awscdk.StackProps
}

func NewPipelineStack(scope constructs.Construct, id string, props *PipelineStackProps) awscdk.Stack{
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	sourceArtifact := awscodepipeline.NewArtifact(jsii.String("source"))
	cloudAssemblyArtifact := awscodepipeline.NewArtifact(jsii.String("cloudassembly"))
	githubSource := awscodepipelineactions.NewCodeStarConnectionsSourceAction(&awscodepipelineactions.CodeStarConnectionsSourceActionProps{
		ActionName: jsii.String("github"),
		Owner: jsii.String("dihmuzikien"),
		Repo: jsii.String("go-cdk-pipeline"),
		Branch: jsii.String("main"),
		Output: sourceArtifact,
		CodeBuildCloneOutput: jsii.Bool(true),
		ConnectionArn: awsssm.StringParameter_ValueForStringParameter(stack, jsii.String("GITHUB_CONNECTOR_ARN"), nil),
	})

	pipeline := pipelines.NewCdkPipeline(stack, jsii.String("cicdPipeline"), &pipelines.CdkPipelineProps{
		CloudAssemblyArtifact: cloudAssemblyArtifact,
		SourceAction: githubSource,
		SelfMutating: jsii.Bool(true),
		CrossAccountKeys: jsii.Bool(false),
		SynthAction: pipelines.NewSimpleSynthAction(&pipelines.SimpleSynthActionProps{
			CloudAssemblyArtifact: cloudAssemblyArtifact,
			SourceArtifact: sourceArtifact,
			InstallCommands: jsii.Strings(
				"make install",
				"cd $HOME/.goenv && git pull --ff-only && cd -",
				"goenv install 1.16.4",
				"goenv local 1.16.4"),
			TestCommands: jsii.Strings("make test"),
			BuildCommands: jsii.Strings("make build"),
			SynthCommand: jsii.String("make synth"),
			EnvironmentVariables: buildEnvVars(),
		}),
	})

	devStage := myapp.NewAppStage(stack, "develop", &myapp.AppStageProps{
		StageProps: awscdk.StageProps{
			Env: &awscdk.Environment{
				Account: stack.Account(),
				Region: jsii.String("us-west-2"),
			},
		},
	})
	pipeline.AddApplicationStage(devStage, nil)

	// prodStage := myapp.NewAppStage(stack, "prod", &myapp.AppStageProps{
	// 	StageProps: awscdk.StageProps{
	// 		Env: &awscdk.Environment{
	// 			Account: stack.Account(),
	// 			Region: jsii.String("us-west-2"),
	// 		},
	// 	},
	// })
	// pipeline.AddApplicationStage(prodStage, &pipelines.AddStageOptions{
	// 	ManualApprovals: jsii.Bool(true),
	// })
	return stack
}

func buildEnvVars() *map[string]*awscodebuild.BuildEnvironmentVariable{
	m := make(map[string]*awscodebuild.BuildEnvironmentVariable)
	m["GO111MODULE"] = &awscodebuild.BuildEnvironmentVariable{
		Value: "on",
		Type: awscodebuild.BuildEnvironmentVariableType_PLAINTEXT,
	}
	return &m
}
