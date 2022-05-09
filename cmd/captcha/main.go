package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type Event struct {
	Token string `json:"token"`
	IP string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Address string `json:"address"`
	NumAvatars int `json:"num_avatars"`
	TransactionNumber int `json:"transaction_number"`
	Nonce string `json:"nonce"`
}




var rdsEndpoint, rdsPwd, secretKey, gURL, siteKey, projectID, contractAddress, signingKey, abi, maxCount, maxIncr, validOrigins, envErr = GetEnvVars()
var rds = NewRedis(rdsEndpoint, rdsPwd, "COUNT")


func HandleRequest(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var event Event
	fmt.Println(e.Headers)
	response := NewErrResponse(GetOriginFromHeaders(e.Headers), validOrigins)
	eventUnmarshalErr := json.Unmarshal([]byte(e.Body), &event)
	if eventUnmarshalErr != nil {
		return response, eventUnmarshalErr
	}
	fmt.Printf("Context: %+v\n\nEvent from Body: %+v\n", ctx, event)
	if rdsVerifyErr := rds.VerifyValidAddress(ctx, event.Address); rdsVerifyErr != nil {
		return response, ConstructErrorResponse(&response, rdsVerifyErr)
	}
	fmt.Println("Verified address")

	if rdsCountCheckErr := rds.GetReservedCount(ctx, event.NumAvatars, maxCount); rdsCountCheckErr != nil {
		return response, ConstructErrorResponse(&response, rdsCountCheckErr)
	}

	fmt.Println("Reserved Count has not reached the max")

	if event.NumAvatars > maxIncr || event.NumAvatars < 1 {
		return response, errors.New("invalid mint amount")
	}

	fmt.Println("Valid amount of NFTs were selected")

	var nonceErr error
	event.Nonce, nonceErr = GetNonce()
	if nonceErr != nil {
		return response, nonceErr
	}

	fmt.Println("Nonce generated")

	googleResponse, googleErr:= VerifyCaptcha(gURL, projectID, secretKey, siteKey, event.Token, event.IP, event.UserAgent, e.Headers["Referer"])
	if googleErr != nil {
		return response, googleErr
	}
	fmt.Println("No Error from Google Captcha Verification")
	if !googleResponse.Token.Valid {
		fmt.Printf("%+v\n", googleResponse)
		return response, errors.New("begone bot")
	}
	fmt.Println("User is not a bot")

	hash, signature, signingErr := SignTxn(event, signingKey)
	if signingErr != nil {
		return response, signingErr
	}
	fmt.Println("Transaction signed")
	queueNum, rdsKeysErr := rds.GetQueueNum(ctx)
	if rdsKeysErr != nil {
		return response, rdsKeysErr
	}
	fmt.Println("Queue number retrieved")

	signedRequest := NewSignedRequest(event, signature, hash, contractAddress, abi, queueNum + 1)

	token, encodingErr := signedRequest.Gzip()
	if encodingErr != nil {
		return response, encodingErr
	}
	fmt.Println("Encoded token")

	fmt.Printf("Address: %s\nToken: %s\n", event.Address, token)
	if rdsSetErr := rds.MarkAddressAsUsed(ctx, event.Address, token); rdsSetErr != nil {
		return response, rdsSetErr
	}
	fmt.Println("Marked Address as used")

	if rdsIncrErr := rds.IncrementCounter(ctx, event.NumAvatars, maxCount); rdsIncrErr != nil {
		return response, rdsIncrErr
	}
	fmt.Println("Incremented Counter")

	return NewResponse(GetOriginFromHeaders(e.Headers), validOrigins, queueNum)
}

func main() {
	fmt.Println("Deploying Captcha lambda!")
	if envErr != nil {
		log.Fatalf("ERROR PARSING ENVIRONMENT VARIABLES: %s\n", envErr.Error())
	}
	defer rds.Close()
	ping, rdsErr := rds.Ping()
	if rdsErr != nil {
		log.Fatalf("ERROR CONNECTING TO REDIS: %s", rdsErr.Error())
	}
	fmt.Println("Redis Ping: " + ping)

	lambda.Start(HandleRequest)
}