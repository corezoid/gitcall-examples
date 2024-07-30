# gitcall-custom-dockerfile

```bash
cd ./js

docker build -t gitcall-example .

docker run --rm -it -p 9999:9999 -e GITCALL_PORT=9999 --user 501:501 --read-only gitcall-example

curl http://127.0.0.1:9999 -H 'Content-Type: application/json' -d '{"jsonrpc":"2.0","id":"task-id-1","method":"Usercode.Run","params":{"key1":"val1"}}'
```