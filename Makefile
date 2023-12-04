build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main src/main.go
	cp src/.env bin/.env

deploy: build
	serverless deploy --stage prod

clean:
	rm -rf ./bin ./vendor Gopkg.lock ./serverless