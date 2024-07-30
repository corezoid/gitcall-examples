package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type request struct {
	Jsonrpc string         `json:"jsonrpc"`
	ID      string         `json:"id"`
	Method  string         `json:"method"`
	Params  map[string]any `json:"params"`
}

type response struct {
	Jsonrpc string         `json:"jsonrpc"`
	ID      string         `json:"id,omitempty"`
	Result  map[string]any `json:"result,omitempty"`
	Error   *respError     `json:"error,omitempty"`
}

type respError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("GITCALL_PORT")
	if port == "" {
		log.Fatal("GITCALL_PORT env is required but not set")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handler)

	fmt.Println("Listening on http://0.0.0.0:" + port)
	http.ListenAndServe("0.0.0.0:"+port, mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer check_panic(w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		send_err(w, "", 1, err.Error())
		return
	}

	var req request
	if err := json.Unmarshal(body, &req); err != nil {
		send_err(w, "", 1, err.Error())
		return
	}
	result, err := usercode(req.Params)
	if err != nil {
		send_err(w, req.ID, 1, err.Error())
		return
	}

	send_ok(w, req.ID, result)
}

// usercode logic
func usercode(data map[string]any) (map[string]any, error) {
	data["golang"] = "Hello, world!"
	return data, nil
}

func send_ok(w http.ResponseWriter, id string, result map[string]any) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(response{
		Jsonrpc: "2.0",
		ID:      id,
		Result:  result,
	})
	if err == nil {
		// nolint
		w.Write(b)
	}
}

func send_err(w http.ResponseWriter, id string, code int, message string) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(response{
		Jsonrpc: "2.0",
		ID:      id,
		Error: &respError{
			Code:    code,
			Message: message,
		},
	})
	if err == nil {
		// nolint
		w.Write(b)
	}
}

func check_panic(w http.ResponseWriter) {
	if r := recover(); r != nil {
		slog.Error("logPanic", "r", r)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		b, err := json.Marshal(response{
			Jsonrpc: "2.0",
			Error: &respError{
				Code:    2,
				Message: fmt.Sprintf("%v", r),
			},
		})
		if err == nil {
			// nolint
			w.Write(b)
		}
	}

}
