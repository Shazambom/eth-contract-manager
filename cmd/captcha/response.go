package main

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
)


type Response struct {
	Token string `json:"token"`
	QueueNumber int64 `json:"queue_number"`
}

func NewResponse(origin string, originList []string, queueNumber int64) (events.APIGatewayProxyResponse, error) {
	resp := Response{
		Token:       "",
		QueueNumber: queueNumber,
	}
	data, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		return events.APIGatewayProxyResponse{}, marshalErr
	}
	if StrInStrList(origin, originList) {
		response := events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      origin,
				"Access-Control-Allow-Credentials": "true",
			},
			StatusCode: 200,
			Body:       string(data),
		}
		return response, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 500}, errors.New("unknown origin")
}

func NewErrResponse(origin string, originList []string) events.APIGatewayProxyResponse {
	if StrInStrList(origin, originList) {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      origin,
				"Access-Control-Allow-Credentials": "true",
			},
			StatusCode: 400,
		}
	}
	return events.APIGatewayProxyResponse{StatusCode: 500}
}

func GetOriginFromHeaders(headers map[string]string) string {
	if val, ok := headers["Origin"]; ok {
		return val
	} else if lilO, lilOk := headers["origin"]; lilOk {
		return lilO
	}
	return ""
}