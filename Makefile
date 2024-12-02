format:
	gofmt -s -w .

clean:
	@go clean
	@rm -rf ./bin

build: clean
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/restapi/bootstrap restapi/main.go
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/websocket/bootstrap websocket/main.go
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/bookingCreated/bootstrap events/bookingCreated/main.go
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/bookingPaid/bootstrap events/bookingPaid/main.go

zip: build
	@zip -j -9 ./bin/restapi.zip ./bin/restapi/bootstrap
	@zip -j -9 ./bin/websocket.zip ./bin/websocket/bootstrap
	@zip -j -9 ./bin/bookingCreated.zip ./bin/bookingCreated/bootstrap
	@zip -j -9 ./bin/bookingPaid.zip ./bin/bookingPaid/bootstrap

deploy-local: zip
	sls deploy --stage local
