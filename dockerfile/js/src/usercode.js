import * as crypto from "node:crypto";

export default async (data) => {  
    data.type = "ESM";
    data.id = crypto.randomBytes(20).toString('hex');
    return data;
};
