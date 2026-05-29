FROM alpine:3.23 as builder

RUN apk add git
WORKDIR /uWebSockets
RUN git clone --recursive https://github.com/uNetworking/uWebSockets --depth=1 --shallow-submodules .

RUN apk add g++ make zlib-dev
RUN make -C uSockets
RUN make install

WORKDIR /nlohmann
RUN wget https://github.com/nlohmann/json/releases/download/v3.12.0/json.hpp

WORKDIR /app
COPY src src

RUN g++ -O3 -Wpedantic -Wall -Wextra -Wsign-conversion -Wconversion -std=c++2a -I/ -I/uWebSockets/uSockets/src/  -flto /app/src/main.cpp  /uWebSockets/uSockets/*.o -lz -o Client

ENTRYPOINT ["./Client"]