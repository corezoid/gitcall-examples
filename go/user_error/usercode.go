package main

import (
	"context"
	"errors"

	"github.com/corezoid/gitcall-go-runner/gitcall"
)

func usercode(ctx context.Context, data map[string]any) error {
	_, ok := data["panic"]
	if ok {
		panic("my custom user panic")
	}
	return errors.New("my custom user error")
}
func main() {
	gitcall.Handle(usercode)
}
