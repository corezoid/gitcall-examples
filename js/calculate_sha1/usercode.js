
// build command: npm install crypto-js

const sha1 = require("crypto-js/sha1");

module.exports = (data) => {
    data.res = sha1(data.in).toString(); 
    
    return data;
};
