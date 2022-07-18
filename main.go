package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/wiredworm/slack-handler/cmd"
)

func main() {
	lambda.Start(cmd.Handler)
}
