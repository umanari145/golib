package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() {
	fmt.Println("hello world")
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
