#include <uWebSockets/src/App.h>
#include <nlohmann/json.hpp>
using json = nlohmann::json;

int main(int argc, char *argv[])
{
	uWS::App().post("/*", [](auto *res, auto *req) {

        std::cout << "Request method is: " << req->getMethod() << "\n";

        const uint64_t contentLength = std::stoi(std::string(req->getHeader("content-length")));
        std::cout << ">: " << contentLength << " ";
        if(contentLength == 0) {
            res->end("'/' request needs a payload");
            std::cout << "contentLength == 0" << std::endl;
            return;
        }

        /* You also need to set onAborted if receiving data */
		res->onData([res, bodyBuffer = (std::string *)nullptr](std::string_view chunk, bool isLast) mutable {
			if (isLast) {
				std::cout << "Sending back posted body now" << std::endl;
				if (bodyBuffer) {
					/* Send back the (chunked) body we got, as response */
					bodyBuffer->append(chunk);
					res->end(*bodyBuffer);
                    std::cout << "[RES] <: " << bodyBuffer << std::endl;
                    json jsonData = json::parse(*bodyBuffer);
                    std::cout << "[jsonrpc1] <: " << jsonData["jsonrpc"] << std::endl;
					delete bodyBuffer;
				} else {
					/* Send back the body we got, as response (fast path) */
					res->end(chunk);
                    std::cout << "[RES] <: " << chunk << std::endl;
                    json jsonData = json::parse(chunk);
                    std::cout << "[jsonrpc2] <: " << jsonData["jsonrpc"] << std::endl;
				}
			} else {
				/* Slow path */
				if (!bodyBuffer) {
					bodyBuffer = new std::string;
				}
				/* If we got the body in a chunk, buffer it up until whole */
				bodyBuffer->append(chunk);
			}
		});

        res->onAborted([=]() {
            std::cout << "HTTP socket was closed!" << std::endl;
        });

	}).listen(9999, [](auto *listenSocket){
        if (listenSocket)
        {
            std::cout << "Listening for connections...\n";
        }
        
    }).run();

    std::cout << "Shoot! We failed to listen and the App fell through, exiting now\n!";

    return 0;
}