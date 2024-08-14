struct Response: Encodable {
     let id: String
     let jsonrpc: String
     let result: [String: Any]?
     let error: ResponseError?
     enum CustomerKeys: String, CodingKey
     {
          case id, jsonrpc, result, error
     }
     func encode (to encoder: Encoder) throws
     {
          var container = encoder.container (keyedBy: CustomerKeys.self)
          try? container.encodeIfPresent (id, forKey: .id)
          try? container.encodeIfPresent (jsonrpc, forKey: .jsonrpc)
          try? container.encodeIfPresent (result, forKey: .result)
          try? container.encodeIfPresent (error, forKey: .error)
     }
}

struct ResponseError: Codable {
    let code: Int
    let message: String
}

struct Request: Decodable {
     let id: String
     let jsonrpc: String
     let params: [String: Any]
     enum CustomerKeys: String, CodingKey
     {
          case id, jsonrpc, params
     }
     init (from decoder: Decoder) throws {
          let container =  try decoder.container (keyedBy: CustomerKeys.self)
          id = try container.decode (String.self, forKey: .id)
          jsonrpc = try container.decode (String.self, forKey: .jsonrpc)
          params = try container.decode ([String: Any].self, forKey: .params)
     }
}