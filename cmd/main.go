package main

import (
	"os"

	"github.com/JAlayon/go-serverless-example/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynamoDBClient dynamodbiface.DynamoDBAPI
)

const tableName = "LambdaInGoUser"

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		panic("An error has occurred during session creation (aws)")
	}
	dynamoDBClient = dynamodb.New(awsSession)
	lambda.Start(handler)

}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, tableName, dynamoDBClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynamoDBClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynamoDBClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynamoDBClient)
	default:
		return handlers.UnhandleMethod()
	}

}
