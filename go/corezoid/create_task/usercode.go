package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/corezoid/gitcall-go-runner/gitcall"
	"github.com/corezoid/sdk-go/corezoid"
)

func usercode(_ context.Context, data map[string]interface{}) error {
	secret, ok := data["api_secret"].(string)
	if !ok {
		return fmt.Errorf("data misses api_secret key")
	}
	key, ok := data["api_key"].(float64)
	if !ok {
		return fmt.Errorf("data misses api_key key")
	}
	convID, ok := data["conv_id"].(float64)
	if !ok {
		return fmt.Errorf("data misses conv_id key")
	}

	auth := &corezoid.ApiKeyAuth{Login: int(key), Secret: secret}
	client := corezoid.New("https://api.corezoid.com", http.DefaultClient)

	ops := corezoid.Ops{}
	ops.Add(corezoid.MapOp{
		"type":    "create",
		"conv_id": convID,
		"obj":     "task",
		"data": map[string]interface{}{
			"test": 123,
		},
	})

	res := client.CallJson(ops, auth).Decode()
	if res.Err != nil {
		return res.Err
	} else if res.RequestProc != "ok" {
		return fmt.Errorf("response: request_proc: %v", res.RequestProc)
	}

	return nil
}

func main() {
	gitcall.Handle(usercode)
}
