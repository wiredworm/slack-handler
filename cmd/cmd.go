package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

var animals = []string{"Donkey", "Monkey", "Lion", "Tiger", "Bear", "Elephant"}

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

	log.Println("Handler Started")
	log.Println("JSON Request:")
	log.Println(JSONify(request))

	if request.HTTPMethod == "GET" {

		log.Println("GET Method Called")

		if request.Resource == "animal" {

			log.Println("GET called on animal resource")

			selectedAnimal := animals[rand.Intn(len(animals))]

			return events.APIGatewayProxyResponse{
				Body:       selectedAnimal,
				StatusCode: 200,
			}, nil
		}
	}

	payload, err := getRequestPayload(request)
	_, _ = payload, err

	return events.APIGatewayProxyResponse{
		Body: Response{
			ReplaceOriginal: true,
			Text:            "Success",
		}.ToString(),
		StatusCode: 200,
	}, nil

}
