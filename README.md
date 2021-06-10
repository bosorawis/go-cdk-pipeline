# Serverless application with CI/CD in Go

## Installation list

- [nodeJS](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [go1.16.5](https://golang.org/doc/install)
- [aws-cli](https://aws.amazon.com/cli/)

## NOTES
This demo only supports **single region/account** deployment. 
The Pipeline construct **CAN** support multi-region multi-account pipeline but that requires additional setup and resources including KMS key which does not have a free tier.

## Structure
```
.
├── infastructure           # cdk code defining all our infrastructure written in go
├── goapp                   # application code for lambda serverless 
```
## Preparing 

### Install and Configure AWS CDK

#### Install
```bash
npm install -g aws-cdk
cdk --version
```

#### Bootstrap your AWS account
```bash
cdk bootstrap 
```
more information about `cdk bootstrap` [here](https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html)

#### Create Github connection for CodePipeline
CodePipeline will need a way to access your Github account. You can follow the official guide [here](https://docs.aws.amazon.com/dtconsole/latest/userguide/connections-update.html) or
##### To create Github connection:
1. Browse to AWS Developer console at https://console.aws.amazon.com/codesuite/settings/connections
2. Go to **Settings > Connections**
3. Click **Create Connection** and follow the prompt

## Deploying the pipeline and application

Clone this repo and edit the following file to use your Github information

```go
// infrastructure/pipeline.go
    githubSource := awscodepipelineactions.NewCodeStarConnectionsSourceAction(&awscodepipelineactions.CodeStarConnectionsSourceActionProps{
		ActionName: jsii.String("github"),
		Owner: jsii.String("<your-github-account>"),
		Repo: jsii.String("<your-github-repo>"),
		Branch: jsii.String("<branch>"),
		Output: sourceArtifact,
		CodeBuildCloneOutput: jsii.Bool(true),
		ConnectionArn: awsssm.StringParameter_ValueForStringParameter(stack, jsii.String("GITHUB_CONNECTOR_ARN"), nil),
	})
```


### First Deployment
```bash
make app
cdk synth
cdk deploy
```
This will create a self-mutating CodePipeline that continuously deploy to `development` environment. 
After this initial deploy, any subsequent deployments will happen automatically - including additional/removal of environments.

More environment can be added by adding following code to `./infrastructure/pipeline.go`

```go
// ./infrastructure/pipeline.go
.
.
.
myStage := myapp.NewAppStage(stack, "stage", &myapp.AppStageProps{
    StageProps: awscdk.StageProps{
        Env: &awscdk.Environment{
            Account: "<your-account-number",
            Region: jsii.String("<your-region>"),
        },
    },
})
pipeline.AddApplicationStage(myStage, nil)
// OR if you need manual approval before deployment
// pipeline.AddApplicationStage(myStage, &pipelines.AddStageOptions{
//    ManualApprovals: jsii.Bool(true),
// })

return stack
```

Then simply do `git push`

The pipeline will mutate itself to add the new stage, then deploy to it all automatically.


