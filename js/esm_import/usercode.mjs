import * as crypto from "node:crypto";

export default async (data) => {  
    data.module = "ESM";
    data.random = crypto.randomBytes(20).toString('hex');
    return data;
};