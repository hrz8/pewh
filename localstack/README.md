## localstack classic mode
Note: execute commands from root dir of the project, not from the localstack dir

### build binary
```bash
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/localstack/bootstrap localstack/main.go && zip -j -9 ./bin/localstack.zip ./bin/localstack/bootstrap
```

### deploy new lambda
```bash
aws lambda create-function --endpoint-url http://localhost:4566 --profile localstack --function-name localstack-test --zip-file fileb://bin/localstack.zip --handler bootstrap --runtime provided.al2 --architectures arm64 --role "arn:aws:iam::000000000000:role/dummy-role" --timeout 60
```

### update lambda
```bash
aws lambda update-function-code --endpoint-url http://localhost:4566 --profile localstack --function-name localstack-test --zip-file fileb://bin/localstack.zip
```

### invoke lambda
```bash
aws lambda invoke --endpoint-url http://localhost:4566 --profile localstack --function-name localstack-test --payload fileb://localstack/input.json localstack/out.json         
```
