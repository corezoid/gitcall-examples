package com.corezoid.gitcall.runner;

import com.sun.net.httpserver.HttpServer;
import java.net.InetSocketAddress;
import java.util.Map;
import java.util.concurrent.Executors;

import com.google.gson.Gson;

import java.nio.charset.StandardCharsets;

public class Gitcall {

    public static void main(String[] args) throws Exception {
        String portStr = System.getenv("GITCALL_PORT");
        if (portStr == null || portStr.isEmpty()) {
            System.out.println("GITCALL_PORT env is required but not set");
            throw new Exception("GITCALL_PORT env is required but not set");
        }

        Integer port = Integer.parseInt(portStr);
        var gson = new Gson();
        var server = HttpServer.create(new InetSocketAddress(port), 0);
        server.setExecutor(Executors.newVirtualThreadPerTaskExecutor());
        server.createContext("/").setHandler(exchange -> {
            String strBody = new String(exchange.getRequestBody().readAllBytes(), StandardCharsets.UTF_8);
            var response = new JsonRpcResponse("2.0", null);
            try {
                var request = gson.fromJson(strBody, JsonRpcRequest.class);
                response.id = request.id;
                System.out.println(String.format("[req] time=%s id=%s", System.currentTimeMillis(), response.id));
                var data = Gitcall.handle(request.params);
                response.result = data;
                System.out.println(String.format("[res] time=%s id=%s", System.currentTimeMillis(), response.id));
            } catch (Exception e) {
                String err = e.toString();
                response.error = new JsonRpcResponse.JsonRpcError(1, err);
                System.out.println(String.format("[res] time=%s id=%s error=%s", System.currentTimeMillis(), response.id, err));
            }
            var jsonResp = gson.toJson(response).getBytes();
            exchange.sendResponseHeaders(200, jsonResp.length);
            try (var os = exchange.getResponseBody()) {
                os.write(jsonResp);
            }
        });
        server.start();
        System.out.println(String.format("Listening on 0.0.0.0:%s", port));
    }

    public static Map<String, Object> handle(Map<String, Object> data) throws Exception {
        data.put("java", "Hello world!");
        return data;
    }

}
