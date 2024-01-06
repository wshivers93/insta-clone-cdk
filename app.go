package main

import (
	"aws-cdk/stacks"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	app := awscdk.NewApp(nil)

	storageStack.StorageStack(app, "StorageStack", awscdk.StackProps {
		Env: &awscdk.Environment {
			Account: jsii.String("388215677591"),
			Region: jsii.String("us-east-1"),
		},
	})

	app.Synth(nil)
	defer jsii.Close()
}
