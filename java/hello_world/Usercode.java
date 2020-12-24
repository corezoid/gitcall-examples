// fq-class name com.corezoid.usercode.Usercode is mandatory
package com.corezoid.usercode;

import com.corezoid.gitcall.runner.api.IUsercode;
import java.util.Map;

public class Usercode implements IUsercode<Map<String, String>, Map<String, String>> {
    @java.lang.Override
    public Map<String, String> handle(Map<String, String> data) throws Exception {
        data.put("hello", "Hello world!");
        return data;
    }
}
