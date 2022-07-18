package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Payload struct {
	Sample string `json:"sample"`
}

type Response struct {
	ResponseType    string `json:"response_type,omitempty"`
	ReplaceOriginal bool   `json:"replace_original"`
	Text            string `json:"text"`
}

func (a Response) ToString() string {
	return JSONify(a)
}

func getRequestPayload(request events.APIGatewayProxyRequest) (Payload, error) {

	// this will umarshall the body the of the request into a struct
	jsonStr, err := url.QueryUnescape(request.Body)
	if err != nil {
		return Payload{}, fmt.Errorf("failed to url decode request body: %v", err)
	}
	byteString := []byte(strings.Trim(jsonStr, "payload="))

	payload := Payload{}
	err = json.Unmarshal(byteString, &payload)
	if err != nil {
		return Payload{}, fmt.Errorf("failed to unmarshall request body: %v", err)
	}

	return payload, nil
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println(JSONify(request))

	payload, err := getRequestPayload(request)
	_ = err

	return events.APIGatewayProxyResponse{
		Body: Response{
			ReplaceOriginal: true,
			Text:            payload.Sample,
		}.ToString(),
		StatusCode: 200,
	}, nil

}
