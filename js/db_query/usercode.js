// build command: npm install pq

const { Client } = require('pg');

module.exports = (data) => {
  return new Promise((resolve, reject) => {
    if(!data.url){
      reject("url is empty");
      return;
    }
    const client = new Client({
      connectionString: data.url,
    });
    client.connect();

    client.query(data.query, (err, res) => {
      data.err = err;
      data.res = res;
      client.end();
      resolve(data);
    });
  });
};
