clean:
	@go clean
	@rm -rf ./bin

build: clean
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/restapi restapi/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/bookingCreated events/bookingCreated/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/bookingPaid events/bookingPaid/main.go

deploy-local: build
	sls deploy -s local --force

zip: build
	@zip -j -9 ./bin/restapi.zip ./bin/restapi
	@zip -j -9 ./bin/bookingCreated.zip ./bin/bookingCreated
	@zip -j -9 ./bin/bookingPaid.zip ./bin/bookingPaid

format:
	gofmt -s -w .
