from fastapi import FastAPI, Body
from usercode import handle

app = FastAPI()

@app.post("/")
def hello(req = Body()):
    jsonrpc = req["jsonrpc"]
    id = req["id"]
    params = req["params"]
    try:
        result = handle(params)
        return {"jsonrpc": jsonrpc, "id": id, "result": result}
    except Exception as err:
        return {"jsonrpc": jsonrpc, "id": id, "error": {"code": 1, "message": str(err)}}