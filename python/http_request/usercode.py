import urllib.request
import json

def handle(data):
    req = urllib.request.Request("https://reqres.in/api/users?page=1", method="GET")
    req.add_header("User-Agent", "urllib-example/0.1")

    f = urllib.request.urlopen(req, timeout=3)

    data["res"] = {
        "code": f.getcode(),
        "body": json.loads(f.read()),
    }

    return data
