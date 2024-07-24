package main

import (
	"context"

	"github.com/corezoid/gitcall-go-runner/gitcall"
	"github.com/google/uuid"
)

func usercode(_ context.Context, data map[string]any) error {
	data["uuid"] = uuid.NewString()
	return nil
}

func main() {
	gitcall.Handle(usercode)
}
