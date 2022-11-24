package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/ricardo-comar/identity-provider/entity"
	"github.com/ricardo-comar/identity-provider/gateway"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		func(o *config.LoadOptions) error {
			o.Region = os.Getenv("AWS_REGION")
			return nil
		})

	if request.HTTPMethod != http.MethodGet {
		return events.APIGatewayProxyResponse{Body: http.StatusText(http.StatusMethodNotAllowed), StatusCode: http.StatusMethodNotAllowed}, err
	}

	var previous string = request.QueryStringParameters["previous"]

	var size int64 = 50
	if len(request.QueryStringParameters["size"]) > 0 {
		size, err = strconv.ParseInt(request.QueryStringParameters["size"], 10, 64)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: "Invalid parameter \"size\", must be numeric", StatusCode: http.StatusBadRequest}, err
		}
	}

	employees, err := gateway.QueryEmployee(cfg, previous, int(size+1))

	if err != nil {
		return events.APIGatewayProxyResponse{Body: http.StatusText(http.StatusInternalServerError), StatusCode: http.StatusInternalServerError}, err
	} else if len(employees) == 0 {
		body, _ := json.Marshal(struct {
			Data []entity.EmployeeData `json:"data"`
		}{Data: []entity.EmployeeData{}})
		return events.APIGatewayProxyResponse{Body: string(body), StatusCode: http.StatusNoContent}, nil
	}

	response := struct {
		Data []entity.EmployeeData `json:"data"`
	}{Data: []entity.EmployeeData{}}

	for _, record := range employees {
		response.Data = append(response.Data, record.CadastroFuncionario)
	}

	var statusCode = http.StatusOK
	if len(response.Data) > int(size) {
		statusCode = http.StatusPartialContent
		response.Data = response.Data[:len(response.Data)-1]
	}

	var body []byte
	body, err = json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: http.StatusText(http.StatusInternalServerError), StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: statusCode}, nil
}
