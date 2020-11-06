package main

import (
	"context"
	"fmt"
	"net/http"

	"encoding/json"

	"bytes"

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

	body, err := json.Marshal(map[string]interface{}{
		"test": 123,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://example.com", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	auth := &corezoid.ApiKeyAuth{Login: int(key), Secret: secret}
	if err := auth.Sign(req); err != nil {
		return err
	}

	data["signed_url"] = req.URL.String()
	data["signed_body"] = string(body)

	return nil
}

func main() {
	gitcall.Handle(usercode)
}
