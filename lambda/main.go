package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/ajaxe/route53updater/cli/shared"
	"github.com/ajaxe/route53updater/lambda/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(r.Body) == 0 {
		return handleResponse(400, "Missing request payload"), nil
	}
	data := shared.Payload{}
	err := json.Unmarshal([]byte(r.Body), &data)
	if err != nil {
		return handleErrorResponse(err), nil
	}
	err = services.ValidateRequest(&data, getPSK())
	if err != nil {
		return handleErrorResponse(err), nil
	}
	svc := services.NewUpdaterService()
	svc.Update(data.IP)
	res := events.APIGatewayProxyResponse{}
	res.StatusCode = 200
	return res, nil
}

func handleErrorResponse(err error) events.APIGatewayProxyResponse {
	log.Printf("error response: %v", err)
	return handleResponse(500, "Internal Server Error")
}

func handleResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}

func getPSK() string {
	if psk, ok := os.LookupEnv("SHARED_KEY"); ok {
		return psk
	}
	return ""
}

func main() {
	lambda.Start(HandleRequest)
}
