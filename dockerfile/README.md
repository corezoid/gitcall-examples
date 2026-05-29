# gitcall-examples / dockerfile

Custom-image examples for Corezoid **GitCall**. Each subdirectory is a self-contained build: a `Dockerfile`, a minimal HTTP/JSON-RPC server, and a `usercode` handler. Use these when the language is not covered by a built-in GitCall runtime, when you need extra system packages, or when you want full control over the image.

Languages: `js`, `python`, `go`, `java`, `php`, `rust`, `cpp`, `swift`, `clojure`, `lisp`, `prolog`.

## Contract

- Server listens on `0.0.0.0:$GITCALL_PORT` (port supplied by GitCall at runtime).
- Accepts `POST /` with a JSON-RPC body: `{"jsonrpc":"2.0","id":"...","method":"Usercode.Run","params":{...}}`.
- Replies with `{"jsonrpc":"2.0","id":"...","result":{...}}` on success, or `{"error":{"code":1,"message":"..."}}` on failure.
- Runs as a non-root user (UID/GID `501`) on a read-only root filesystem.

Replace the `usercode` handler in `src/` (or the language equivalent) with your own logic; keep the server scaffolding as-is.

## Connect to a GitCall node

Point the GitCall node at this repo (or your fork), set the path to one of the subdirectories (e.g. `dockerfile/python`), and pick **Custom Dockerfile** as the runtime. Corezoid will build the image and route tasks to it.

## Build and test locally

```bash
cd ./js   # or python, go, rust, ...

docker build -t gitcall-example .

docker run --rm -it -p 9999:9999 -e GITCALL_PORT=9999 --user 501:501 --read-only gitcall-example

curl http://127.0.0.1:9999 \
  -H 'Content-Type: application/json' \
  -d '{"jsonrpc":"2.0","id":"task-id-1","method":"Usercode.Run","params":{"key1":"val1"}}'
```

Expected response (for the bundled `hello_world` handler):

```json
{"jsonrpc":"2.0","id":"task-id-1","result":{"key1":"val1","hello":"Hello World!"}}
```
