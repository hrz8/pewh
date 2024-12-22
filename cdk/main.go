package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewStage(app, "Local", &StageProps{
		StageProps: awscdk.StageProps{
			Env: localEnv(),
		},
	})

	NewStage(app, "Dev", &StageProps{
		StageProps: awscdk.StageProps{
			Env: devEnv(),
		},
	})

	NewStage(app, "Prod", &StageProps{
		StageProps: awscdk.StageProps{
			Env: prodEnv(),
		},
	})

	app.Synth(nil)
}

func localEnv() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("000000000000"),
		Region:  jsii.String("us-east-1"),
	}
}

func devEnv() *awscdk.Environment {
	return nil
}

func prodEnv() *awscdk.Environment {
	return nil
}
