# gitcall-examples

Code examples for [Corezoid](https://corezoid.com/) **GitCall** nodes in multiple languages: JavaScript, Python, Go, Java, PHP, Clojure, Common Lisp, Prolog (plus C++, Rust, Swift in custom Dockerfiles).

GitCall executes your code from a Git repository as a step inside a Corezoid process. Each request is delivered to the handler as a JSON-RPC `Usercode.Run` call; the value you return becomes the task payload of the next node.

## Repository layout

- `js/`, `python/`, `go/`, `java/`, `php/`, `clojure/`, `lisp/`, `prolog/` — handler-only examples (`hello_world`, `http_request`, `user_error`, dependencies, etc.). Use these when GitCall builds the runtime for you.
- `dockerfile/` — full custom-image examples (Dockerfile + HTTP/JSON-RPC server + handler) for advanced cases or languages without a built-in runtime. See [`dockerfile/README.md`](./dockerfile/README.md).

## Connect an example to a GitCall node

1. In Corezoid, open a process and add a **GitCall** node.
2. Set the repository URL (this repo or your fork) and branch.
3. Set the path to the example directory, e.g. `js/hello_world` or `python/http_request`.
4. Pick the matching language/runtime (or **Custom Dockerfile** when using anything under `dockerfile/`).
5. Save — Corezoid will pull, build, and invoke the handler on each task entering the node.

The handler receives the task payload as its argument and must return the next payload (or throw to produce an error).

## Run and test locally

Handler-only examples can be wrapped with the matching image under `dockerfile/<lang>/` to reproduce the GitCall runtime on your machine. To verify a custom-image example end-to-end:

```bash
cd dockerfile/js   # or python, go, rust, ...

docker build -t gitcall-example .
docker run --rm -it -p 9999:9999 -e GITCALL_PORT=9999 --user 501:501 --read-only gitcall-example

curl http://127.0.0.1:9999 \
  -H 'Content-Type: application/json' \
  -d '{"jsonrpc":"2.0","id":"task-1","method":"Usercode.Run","params":{"key":"value"}}'
```

A successful response has the shape `{"jsonrpc":"2.0","id":"task-1","result":{...}}`; failures return an `error` object with `code` and `message`.

## Docs

- Corezoid: <https://corezoid.com/>
- GitCall reference: see the official Corezoid documentation.
