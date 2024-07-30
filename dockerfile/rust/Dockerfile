FROM rust:1.80-alpine3.20 as builder
RUN apk add --no-cache musl-dev

WORKDIR /usr/src/app

COPY Cargo.toml ./
COPY src src
RUN cargo build --release

FROM alpine:3.20
COPY --from=builder /usr/src/app/target/release/gitcall-rust-app /usr/local/bin/gitcall-rust-app

RUN addgroup --gid 501 usercode && \
    adduser --disabled-password \
    --gecos "" \
    --shell /usr/sbin/nologin \
    --ingroup usercode \
    --uid 501 \
    usercode
USER usercode

CMD ["gitcall-rust-app"]