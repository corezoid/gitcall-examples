package com.corezoid.gitcall.runner;

import java.util.Map;

class JsonRpcResponse {
    public String jsonrpc;
    public String id;
    public Map<String, Object> result;
    public JsonRpcError error;

    static class JsonRpcError {
        public Integer code;
        public String message;
        JsonRpcError(Integer code, String message) {
            this.code = code;
            this.message = message;
        }
    }

    JsonRpcResponse(String jsonrpc, String id) {
        this.jsonrpc = jsonrpc;
        this.id = id;
    }
}