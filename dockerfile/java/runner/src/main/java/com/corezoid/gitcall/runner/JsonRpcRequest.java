package com.corezoid.gitcall.runner;

import java.util.Map;


class JsonRpcRequest {
    public String jsonrpc;
    public String id;
    public String method;
    public Map<String, Object> params;
}