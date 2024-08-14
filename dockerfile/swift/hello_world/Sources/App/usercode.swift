extension String: Error {}

func usercode(_ data: [String: Any]) throws -> [String: Any] {
    var result = data
    result["swift"] = "Hello World!"
    // throw "My custom error"
    return result
}