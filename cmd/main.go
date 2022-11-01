package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := ApiRequest{}
	err := fastjson.Validate(request.Body)

	if err != nil {
		log.Println("Invalid JSON Format")
		return events.APIGatewayProxyResponse{Body: "Invalid JSON format", StatusCode: 500}, nil
	}

	body := request.Body

	json.Unmarshal([]byte(body), &req)
	message := fmt.Sprintf("Welcome %s %s", req.Fname, req.Lname)

	resp := ApiResponse{true, message}
	content, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{Body: string(content), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
