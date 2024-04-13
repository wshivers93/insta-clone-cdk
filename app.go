package main

import (
	"aws-cdk/stacks"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

type appConfig struct {
	AppAccount string `json:"applicationAccount"`
	Region string `json:"applicationRegion"`
}


func main() {
	file, fileErr := os.ReadFile("./config.json")
	
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	
	config := appConfig{}

	if readErr := json.Unmarshal(file, &config); readErr != nil {
		panic(readErr)
	}
	
	app := awscdk.NewApp(nil)

	storageStack.StorageStack(app, "StorageStack", awscdk.StackProps {
		Env: &awscdk.Environment {
			Account: jsii.String(config.AppAccount),
			Region: jsii.String(config.Region),
		},
	})

	app.Synth(nil)
	defer jsii.Close()
}
