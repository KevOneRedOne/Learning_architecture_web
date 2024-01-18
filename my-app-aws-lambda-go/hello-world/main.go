package main

import (
	"context"
	"fmt"
	"os"

	brevo "github.com/getbrevo/brevo-go/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s!\n", sourceIP)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func brevoSMTP() {
	var ctx context.Context
	cfg := brevo.NewConfiguration()

	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", os.Getenv("BREVO_API_KEY"))

	//Configure API key authorization: partner-key
	// cfg.AddDefaultHeader("partner-key", "YOUR_API_KEY")

	br := brevo.NewAPIClient(cfg)
	result, resp, err := br.AccountApi.GetAccount(ctx)
	if err != nil {
		fmt.Println("Error when calling AccountApi->get_account: ", err.Error())
		return
	}
	fmt.Println("GetAccount Object:", result, " GetAccount Response: ", resp)
	return
}

func main() {
	lambda.Start(handler)
}
