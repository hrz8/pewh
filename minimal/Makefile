format:
	gofmt -s -w .

clean:
	@go clean
	@rm -rf ./bin

build: clean
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/hello/bootstrap hello/main.go

zip: build
	@zip -j -9 ./bin/hello.zip ./bin/hello/bootstrap

deploy-local: zip
	sls deploy --stage local

deploy-local-cdk: zip
	cdklocal bootstrap
	cdklocal deploy "Local/*" --require-approval never --force
