package main

import (
    "context"
    "testing"
)

func TestHandler(t *testing.T) {
    got, err := handler(context.Background())
    if err != nil {
        t.Fatal(err)
    }

    if got != "Hello, s3db!" {
        t.Fatalf("got %q", got)
    }
}
