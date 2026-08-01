package main

import (
    "context"
    "fmt"
    "os"
    "log"

    "github.com/aws/aws-lambda-go/lambda"
    handler "github.com/szabozsoltbors/s3db/internal/adapter/in/lambda"
    commands "github.com/szabozsoltbors/s3db/internal/application/commands"
    repository "github.com/szabozsoltbors/s3db/internal/adapter/out/repository/aws/s3"
    "github.com/aws/aws-sdk-go-v2/config"
    s3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

    ctx := context.Background()

	// Load AWS configuration (region, credentials, etc.)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	// Create the S3 client
	s := s3.NewFromConfig(cfg)

    r := repository.NewRepository(s, "szs-dev-s3db-store")
    c := commands.NewCreateCommand(r)
    d := commands.NewDeleteCommand(r)
    l := commands.NewListCommand(r)
    h := handler.NewHandler(c, d, l)

    if os.Getenv("LOCAL") == "1" {
        result, err := h.Handle(context.Background())
        fmt.Println("result:", result)
        fmt.Println("error:", err)
        return
    }

    lambda.Start(h.Handle)
}
