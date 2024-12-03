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
aws sns publish --endpoint-url http://localhost:4566 --profile localstack \
    --topic-arn arn:aws:sns:us-east-1:000000000000:booking-created-local \
    --message '{"key": "value"}'
aws sns publish --endpoint-url http://localhost:4566 --profile localstack \
    --topic-arn arn:aws:sns:us-east-1:000000000000:booking-paid-local \
    --message '{"key": "value"}'
```

### thru publish sqs
```sh
aws sqs send-message --endpoint-url http://localhost:4566 --profile localstack \
    --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/booking-canceled-local \
    --message-body '{"key": "value", "status": "canceled"}'
# using fifo
aws sqs send-message --endpoint-url http://localhost:4566 --profile localstack \
    --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/check-in-started-local.fifo \
    --message-body '{"key": "value"}' \
    --message-group-id "group-1"
```

### invoke directly to lambda sns
```sh
aws lambda invoke --endpoint-url http://localhost:4566 --profile localstack --function-name pewh-local-bookingCreated --payload fileb://input.json /dev/stdout
aws lambda invoke --endpoint-url http://localhost:4566 --profile localstack --function-name pewh-local-bookingPaid --payload fileb://input.json /dev/stdout
```

aws lambda invoke --endpoint-url http://localhost:4566 --profile localstack --function-name pewh-local-bookingCanceled --payload fileb://input.json /dev/stdout
