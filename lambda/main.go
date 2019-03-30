package main

import (
	"github.com/ajaxe/route53updater/cli/shared"
	"github.com/ajaxe/route53updater/lambda/services"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data := shared.Payload
	err := json.Unmarshal(byte(r.Body), &data)
	if err != nil {

	}
	svc := services.NewUpdaterService()
	svc.
	res := events.APIGatewayProxyResponse{}
	res.StatusCode = 200
	return res, nil
}



func main() {
	lambda.Start(HandleRequest)
}
