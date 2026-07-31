package main

import (
    "context"
    "fmt"
    "os"

    "github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (string, error) {
    return "Hello, s3db!", nil
}

func main() {
    if os.Getenv("LOCAL") == "1" {
        result, err := handler(context.Background())
        fmt.Println("result:", result)
        fmt.Println("error:", err)
        return
    }

    lambda.Start(handler)
}
