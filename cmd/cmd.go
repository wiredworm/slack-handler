package cmd

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	ResponseType    string `json:"response_type,omitempty"`
	ReplaceOriginal bool   `json:"replace_original"`
	Text            string `json:"text"`
}

func (a Response) ToString() string {
	return JSONify(a)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println(JSONify(request))

	return events.APIGatewayProxyResponse{
		Body: Response{
			ReplaceOriginal: true,
			Text:            "R:Tape Loading Error",
		}.ToString(),
		StatusCode: 200,
	}, nil
}
