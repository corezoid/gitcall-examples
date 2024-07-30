const http = require('node:http');

const port = process.env.GITCALL_PORT;
if (!port) {
    console.error('GITCALL_PORT env is required but not set');
    process.exit(1);
}

process.on('SIGTERM', () => process.exit(1));
process.on('SIGINT', () => process.exit(1));

const server = http.createServer((request, response) => {
    if (request.method === 'POST' && request.url === '/') {
        let body = [];
        request
            .on('data', chunk => body.push(chunk))
            .on('end', () => {
                response.statusCode = 200;
                handler(Buffer.concat(body).toString(), response);
            });
    } else {
        response.statusCode = 404;
        response.end();
    }
});

const handler = (body, response) => {
    const req = JSON.parse(body);
    const jsonrpc = req.jsonrpc;
    const id = req.id;
    const params = req.params;
    try {
        if (typeof params !== 'object') {
            throw Error('expected request params is object');
        }
        const result = usercode(id, params);
        response.end(JSON.stringify({
            jsonrpc: jsonrpc,
            id: id,
            result: result,
        }));
    } catch (e) {
        response.end(JSON.stringify({
            jsonrpc: jsonrpc,
            id: id,
            error: {
                code: 1,
                message: e.toString(),
            },
        }));
    }

};

const usercode = (taskId, data) => {
    data["js"] = "Hello, world!"
    return data
};

console.log('Listening on 0.0.0.0:' + port);
server.listen(port);