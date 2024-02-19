package main

import (
	"github.com/Bancar/lambda-go"
	"test-lambda-pkg-aws-lambda/service"
)

func main() {

	lambda.Start(service.New().Service)
}
