package web

import (
	"contract-service/utils"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
)


type Response struct {
	Token string `json:"token"`
	QueueNumber int64 `json:"queue_number"`
}

type ErrorResponse struct {
	Error string `json:"error"`
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
	if utils.StrInStrList(origin, originList) {
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
	if utils.StrInStrList(origin, originList) {
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

func ConstructErrorResponse(response *events.APIGatewayProxyResponse, err error) error {
	errBody, marshErr := json.Marshal(ErrorResponse{err.Error()})
	if marshErr != nil {
		return marshErr
	}
	response.Body = string(errBody)
	return nil
}
