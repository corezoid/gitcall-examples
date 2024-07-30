FROM golang:1.22-alpine as builder

WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -v -o /app/gitcall-app ./main.go
# Compress GoLang App binary. Lower size - x4
RUN apk add --no-cache --virtual .fetch-deps upx
RUN upx -1 /app/gitcall-app


FROM alpine:3.20
COPY --from=builder /app/gitcall-app /app/gitcall-app

RUN addgroup --gid 501 usercode && \
    adduser --disabled-password \
    --gecos "" \
    --shell /usr/sbin/nologin \
    --ingroup usercode \
    --uid 501 \
    usercode
USER usercode

ENTRYPOINT ["/app/gitcall-app"]