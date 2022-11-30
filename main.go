package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"how old are you?"`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func HandleLambdaFunction(e MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", e.Name, e.Age)}, nil
}

func main() {
	log.Println("Starting lamdba function")

	lambda.Start(HandleLambdaFunction)
}
