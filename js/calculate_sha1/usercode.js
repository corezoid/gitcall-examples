const crypto = require("crypto");

module.exports = (data) => {
  const input = String(data?.in ?? "");

  data.res = crypto.createHash("sha1").update(input).digest("hex");

  return data;
};
