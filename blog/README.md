### Build a Go Serverless REST APIs and Deploy to AWS using Serverless Framework

Imagine your app scaling effortlessly and you're very far from the stress of server management. That's how serverless `sounds` like, and Go is the perfect language to lead the charge. This article takes you on a hands-on journey, where you'll leverage the Go's powerful web framework GIN and AWS lamda GIN proxy to build efficient serverless APIs and seamlessly deploy them to AWS Cloud, using the magic of Serverless Framework.

### Leeroy Jenkins! Let's get started 🚀

### Let's setup Gin
To spin up our server, we need Gin. Gin is a powerful web framework written in go. It claims to be 40x faster than Martini, a technology of same genre. 
> If you need blasting performance, get yourself some Gin

- To install gin, go ahead run
```shell
go get github.com/gin-gonic/gin
```
- Now you can setup a minimal Gin router. For that, copy the example in the gin documentation and paste it in the main.go file

```go
// main.go

func main(){
    r := gin.Default()

    r.Get("/ping", (c *gin.Context){
        c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
    })

    r.Run("8000")
}
```

- When you're all good, you can now run the server
```shell
go run main.go
```

- When you visit `localhost:8000/ping`, you should see expected response.


## Let's setup the Serverless framework
### First thing first: What is `Serverless Framework` ?

Serverless is a tool that makes it easy to deploy serverless applications to AWS.
> **Zero-Friction Serverless Apps on AWS Lambda**
> Deploy APIs, scheduled tasks, workflows and event-driven apps to AWS Lambda easily with the Serverless Framework.

It removes the hell related to deployment in so many ways. Yeah! `Zero-friction serverless development` Indeed. Why? Because:

- You define your applications as functions and events by declaring AWS Lambda functions and their triggers through simple abstract syntax in YAML.

- After deploying your serverless app in a single command, Lambda functions, triggers & code will be deployed and wired together in the cloud, automatically.

- You can also extend your use cases by installing thousands of Serverless Framework Plugins.

### Installation of `Serverless`
- To install the serverless framework, we need `nodejs`. If you don't have nodejs installed, you can go ahead and install it from there

```javascript
    // Nodejs offical website link
```
- Simply run the command `npm install -g serverless` to get the serverless framework.
- To make sure it has been wen installed, you can run `serverless --version`. You should have a similar outupt.

### Let's define our serverless application
Now that we have serverless installed, we can go ahead and define our serverless app.
As mention above, we define our application events and their triggers in a simple `YAML` file.
#### Let's do it 🎉

In the root of the project, create a file name `serverless.yaml`
```YAML
service: serverless-go-gin-app
frameworkVersion: '>= 3.38.0'
useDotenv: true
provider:
  name: aws
  runtime: go1.x
  region: ${opt:region}
  stage: ${opt:stage}

functions:
  api:
    memorySize: 3008
    timeout: 10
    handler: bin/main
    events:
      - http:
          path: /ping
          method: GET
      - http:
          path: /ping
          method: POST

package:
  patterns:
   - "!*/**"
```

### Let's break this `YAML` configuration down.


Now we're reading to deploy 🚀

Before that, let's setup a `Makefile` to avoid typing the same commands repetitively
Create a `Makefile` in the root of your project and paste this content.
```shell
build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go

deploy: build
	serverless deploy --stage prod

clean:
	rm -rf ./bin ./vendor Gopkg.lock ./serverless
```
> [NOTE]: You'll need to setup your `AWS` profile before deploying, if it's not yet done.
This can be helpful in setting up your account on the local machine.


After setting up your AWS account, you can go ahead and deploy the serverless app by running this command
```shell
make deploy
```

It will take few minutes to `Serverless` to compile, package and deploy the application. While waiting, stay calm `fingers crossed`. (If you want you can pray for it to work on first try 😂😂😂). Anyways....

When the deployment is complete, a link will be generated for you to test your APIs. You can click on that link and test the `/ping` GET route. Normally you should see a json response like so:
```json
{
    "message": "pong!"
}
```

### But Guess what ? It will never work. Instead, you will receive a timeout error. Something like
```md
Connection timeout error blaaablaaablaaablaaa.
```
If this happens, because there is a reason. Let me explain why.

### Why the handler function times out ?

You are receiving a response like:
```json
{
    "message": "pong!"
}
```
Well! The reason is that we're not properly handling our requests. What happened is that, the server has started, and has emitted the request. In fact, the request in not formatted in a way that AWS requires it to be formatted.

The AWS Documentation for **Lambda function handler in Go** says:

*The Lambda function handler is the method in your function code that processes events. When your function is invoked, Lambda runs the handler method. Your function runs until the handler returns a response, exits, or times out.*
*A Lambda function written in Go is authored as a Go executable. In your Lambda function code, you need to include the `github.com/aws/aws-lambda-go/lambda` package, which implements the Lambda programming model for Go. In addition, you need to implement handler function code and a main() function.*

Indeed, we missed a lot from what the documentation says! Let's fix this.


### Refactoring in AWS style

1. First we need to add the package mentionned in the documentation.
```shell
go get github.com/aws/aws-lambda-go/lambda
```
2. Including Gin proxy handler in main
We now need to refactor our `main.go` file to:
    1. Add the handler function code lamda function
    2. Start the lamdba server to listen to incoming requests
    3. Wires the lamdba function responses to our GIN server: This proxy will act like a bridge between our AWS lamda function and our GIN server

```go
//main.go
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)


// AWS gin lambda adapter
var ginLambda *ginadapter.GinLambda

func init() {
	//Set the router as the default one provided by Gin
	r := gin.Default()

	r.Get("/ping", (c *gin.Context){
        c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
    })

	// Start and run the server
	ginLambda = ginadapter.New(r)
}

// AWS Lambda Proxy Handler
// This handler acts like a bridge between AWS Lambda and our Local GIn server
// It maps each GIN route to a Lambda function as handler
//
// This is useful to make our function execution possible.
func GinRequestHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	// Starts Lambda server
	lambda.Start(GinRequestHandler)

}
```

After updating our `main.go` file, we can save the file and redeploy again.
When the deployment is successful, you should see the list of the endpoints with their URLs.
Go ahead and test the `/ping` GET route. This time you should receive a success response (for real)

```json
{
    "message": "pong!"
}
```
That's it! It worked 🎉

### Wait! What happened under the hood ?
To say all in one, `Serverless` took care of everything for us.
In fact, serverless uses our expressions and configurations in the `serverless.yaml` file to create a `AWS Cloud Formation Stack`.
- After your application is compiled,
- Serverless takes the compilation binary and package it up
- The binary is then uploaded to a `AWS S3 Bucket`
- Then this binary code will be attached to a Lambda function as handler
- The lamda function is triggered `on` an `AWS API Gateway` event

All this has been handled by `Serverless`, based on our configuration file.

### What's next ?
This is a minimal example of creating serverless RESTful APIs using Go and AWS.
You can do few things to enhance the application. For example
- Connecting the applicattion to a Database, `Amazon Dynamo DB`
- Refactoring the codebase
- Connecting a custom domain name to your APIs endpoint

> You can check the full source code for this tutorial in my github
https://github.com/baimamboukar/go-serverless-api

Thanks for reading! 
My name is Baimam Boukar, I'm a Software Engineer and I enjoy sharing my knowledge trough blog posts. I write about Serverless and Cloud native applications in Go, Flutter mobile app development and Bitcoin development. Let's stay connected 
- Find me on Github
- Let's connect on LinkedIn


### References