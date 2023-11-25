build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go

deploy_production: build
	serverless deploy --stage prod