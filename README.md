# pewh
Yeah right here I'm learning to use Serverless Framework with Golang.

## local development
```sh
make deploy-local
```

## invoke restapi

### get apigateway resources
```sh
# get restapi_id
awslocal apigateway get-rest-apis --profile localstack
# get resources
awslocal apigateway get-resources --rest-api-id <restapi_id> --profile localstack
```

### thru apigateway using curl
```sh
curl -X GET \
    -H "Content-Type: application/json" \
    http://localhost:4566/restapis/<restapi_id>/local/_user_request_/api/v1/rooms
curl -X GET \
    -H "Content-Type: application/json" \
    http://localhost:4566/restapis/<restapi_id>/local/_user_request_/api/v1/rooms
```

### invoke directly to lambda
```sh
awslocal lambda invoke --profile localstack \
    --function-name pewh-local-restapi \
    --payload '{"path": "/api/v1/rooms", "httpMethod": "GET", "queryStringParameters": {"key": "value"}}' \
    /dev/stdout
awslocal lambda invoke --profile localstack \
    --function-name pewh-local-restapi \
    --payload '{"path": "/api/v1/rooms/another", "httpMethod": "GET", "queryStringParameters": {"key": "value"}}' \
    /dev/stdout
```

## invoke event handler

### thru publish sns
```sh
awslocal sns publish --profile localstack \  
  --topic-arn arn:aws:sns:us-east-1:000000000000:booking-created-local \
  --message '{"key": "value"}'
awslocal sns publish --profile localstack \
  --topic-arn arn:aws:sns:us-east-1:000000000000:booking-paid-local \
  --message '{"key": "value"}'
```

### invoke directly to lambda
```sh
awslocal lambda invoke --profile localstack \
  --function-name pewh-local-bookingCreated \
  --payload '{
    "Records": [
      {
        "EventSource": "aws:sns",
        "EventVersion": "1.0",
        "EventSubscriptionArn": "arn:aws:sns:us-east-1:000000000000:booking-created-local",
        "Sns": {
          "Type": "Notification",
          "MessageId": "12345678-1234-5678-9012-123456789012",
          "TopicArn": "arn:aws:sns:us-east-1:000000000000:booking-created-local",
          "Subject": null,
          "Message": "{\"key\": \"value\"}",
          "Timestamp": "2024-11-29T10:30:00.000Z",
          "SignatureVersion": "1",
          "Signature": "EXAMPLESIGNATURE",
          "SigningCertUrl": "https://sns.us-east-1.amazonaws.com/SimpleNotificationService.pem",
          "UnsubscribeUrl": "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:000000000000:booking-created-local",
          "MessageAttributes": {}
        }
      }
    ]
  }' /dev/stdout
awslocal lambda invoke --profile localstack \
  --function-name pewh-local-bookingPaid \
  --payload '{
    "Records": [
      {
        "EventSource": "aws:sns",
        "EventVersion": "1.0",
        "EventSubscriptionArn": "arn:aws:sns:us-east-1:000000000000:booking-paid-local",
        "Sns": {
          "Type": "Notification",
          "MessageId": "12345678-1234-5678-9012-123456789012",
          "TopicArn": "arn:aws:sns:us-east-1:000000000000:booking-paid-local",
          "Subject": null,
          "Message": "{\"key\": \"value\"}",
          "Timestamp": "2024-11-29T10:30:00.000Z",
          "SignatureVersion": "1",
          "Signature": "EXAMPLESIGNATURE",
          "SigningCertUrl": "https://sns.us-east-1.amazonaws.com/SimpleNotificationService.pem",
          "UnsubscribeUrl": "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:000000000000:booking-paid-local",
          "MessageAttributes": {}
        }
      }
    ]
  }' /dev/stdout
```

