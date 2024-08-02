import time
import logging
from fastapi import FastAPI, Body
from usercode import handle

logger = logging.getLogger('uvicorn.error')
logger.setLevel(logging.DEBUG)

app = FastAPI()

@app.post("/")
def request(req = Body()):
    jsonrpc = req["jsonrpc"]
    id = req["id"]
    params = req["params"]
    try:
        logger.info("[req] time=%d id=%s" % (round(time.time() * 1000), id))
        result = handle(params)
        logger.info("[res] time=%d id=%s" % (round(time.time() * 1000), id))
        return {"jsonrpc": jsonrpc, "id": id, "result": result}
    except Exception as err:
        logger.info("[res] time=%d id=%s error=%s" % (round(time.time() * 1000), id, str(err)))
        return {"jsonrpc": jsonrpc, "id": id, "error": {"code": 1, "message": str(err)}}