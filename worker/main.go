package main

import (
    "context"
    "fmt"
    "os"

    "github.com/aws/aws-lambda-go/lambda"
    handler "github.com/szabozsoltbors/s3db/internal/adapter/in/lambda"
    commands "github.com/szabozsoltbors/s3db/internal/application/commands"
    repository "github.com/szabozsoltbors/s3db/internal/adapter/out/repository/aws/s3"
)

func main() {
    r := repository.NewRepository()
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
