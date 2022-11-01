package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/valyala/fastjson"
)

type ApiRequest struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

type ApiResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func getResponse(status bool, message string) string {
	msg := fmt.Sprintf(message)
	resp := ApiResponse{status, msg}
	content, _ := json.Marshal(resp)
	return string(content)
}

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")))

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := ApiRequest{}
	err := fastjson.Validate(request.Body)

	if err != nil {
		content := getResponse(false, "JSON Validation error")
		return events.APIGatewayProxyResponse{Body: content, StatusCode: 500}, nil
	}

	body := request.Body

	json.Unmarshal([]byte(body), &req)

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(session)

	result, err := dynamodbattribute.MarshalMap(req)

	if err != nil {
		content := getResponse(false, "Parsing JSON error")
		return events.APIGatewayProxyResponse{Body: content, StatusCode: 500}, nil
	}

	input := &dynamodb.PutItemInput{
		Item:      result,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		content := getResponse(false, "Writing to database error")
		return events.APIGatewayProxyResponse{Body: content, StatusCode: 500}, nil
	}

	content := getResponse(true, fmt.Sprintf("Welcome %s %s", req.Fname, req.Lname))
	return events.APIGatewayProxyResponse{Body: content, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
