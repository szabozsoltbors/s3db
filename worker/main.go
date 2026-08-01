package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	handler "github.com/szabozsoltbors/s3db/internal/adapter/in/lambda"
	repository "github.com/szabozsoltbors/s3db/internal/adapter/out/repository/aws/s3"
	commands "github.com/szabozsoltbors/s3db/internal/application/commands"
)

func main() {
	// Load environment variables from .env file
	var bucket string = os.Getenv("S3_BUCKET_NAME")

	// Load AWS configuration (region, credentials, etc.)
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	// Composition root: create repository, commands, and handler
	s := s3.NewFromConfig(cfg)
	r := repository.NewRepository(s, bucket)
	c := commands.NewCreateCommand(r)
	d := commands.NewDeleteCommand(r)
	l := commands.NewListCommand(r)
	h := handler.NewHandler(c, d, l)

	// If running locally, invoke the handler directly for testing
	if os.Getenv("LOCAL") == "1" {
		result, err := h.Handle(context.Background())
		fmt.Println("result:", result)
		fmt.Println("error:", err)
		return
	}

	// Start the AWS Lambda function
	lambda.Start(h.Handle)
}
