// fq-class name com.corezoid.usercode.Usercode is mandatory
package com.corezoid.usercode;

import com.corezoid.gitcall.runner.api.UsercodeHandler;
import java.util.Map;

public class Usercode implements UsercodeHandler<Map<String, String>, Map<String, String>> {
    @java.lang.Override
    public Map<String, String> handle(Map<String, String> data) throws Exception {
        data.put("hello", "Hello world!");
        return data;
    }
}
