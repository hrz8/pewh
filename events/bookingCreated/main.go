package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(_ context.Context, snsEvent events.SNSEvent) error {
	for _, record := range snsEvent.Records {
		snsMessage := record.SNS.Message
		fmt.Printf("Message ID: %s\n", "N/A")
		fmt.Printf("Received message: %s\n", snsMessage)

		var payload map[string]interface{}
		if err := json.Unmarshal([]byte(snsMessage), &payload); err != nil {
			log.Printf("Failed to parse message: %v\n", err)
			continue
		}

		log.Printf("Payload Booking Create: %+v\n", payload)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
