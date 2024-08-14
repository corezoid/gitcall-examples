import Vapor

func routes(_ app: Application) throws {
    app.post { req async throws -> String in
        let request = try req.content.decode(Request.self)
        req.logger.info("[req] time=\(Date().millisecondsSince1970) id=\(request.id)")
        do {
            let result = try usercode(request.params)
            let resp = Response(id: request.id, jsonrpc: request.jsonrpc, result: result, error: nil)
            let jsonEncoder = JSONEncoder()
            let jsonResultData = try? jsonEncoder.encode(resp)
            req.logger.info("[res] time=\(Date().millisecondsSince1970) id=\(request.id)")
            return String(data: jsonResultData!, encoding: .utf8) ?? ""
        } catch let error {
            let message = "\(error)"
            let resp = Response(id: request.id, jsonrpc: request.jsonrpc, result: nil, error: ResponseError(code: 1, message: message))
            let jsonEncoder = JSONEncoder()
            let jsonResultData = try? jsonEncoder.encode(resp)
            req.logger.info("[res] time=\(Date().millisecondsSince1970) id=\(request.id) error=\(error)")
            return String(data: jsonResultData!, encoding: .utf8) ?? ""
        }    
    }
}

extension Date {
    var millisecondsSince1970:Int64 {
        Int64((self.timeIntervalSince1970 * 1000.0).rounded())
    }    
    init(milliseconds:Int64) {
        self = Date(timeIntervalSince1970: TimeInterval(milliseconds) / 1000)
    }
}