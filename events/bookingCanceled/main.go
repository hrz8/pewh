package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		fmt.Printf("Message ID: %s\n", message.MessageId)
		fmt.Printf("Received message: %s\n", message.Body)

		var payload map[string]interface{}
		if err := json.Unmarshal([]byte(message.Body), &payload); err != nil {
			log.Printf("Failed to parse message: %v\n", err)
			continue
		}

		log.Printf("Payload Booking Canceled: %+v\n", payload)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
