build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go

deploy: build
	serverless deploy --stage prod

clean:
	rm -rf ./bin ./vendor Gopkg.lock ./serverless