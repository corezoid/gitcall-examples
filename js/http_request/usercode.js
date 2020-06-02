const https = require('https');

module.exports = (data) => {
  return new Promise((resolve) => {
    https
      .get('https://reqres.in/api/users?page=1', (resp) => {
        let body = '';
        resp.on('data', (chunk) => {
          body += chunk;
        });
        resp.on('end', () => {
          data.res = JSON.parse(body);
          resolve(data);
        });
      })
      .on('error', (err) => {
        data.err = err;
        resolve(data);
      });
  });
};