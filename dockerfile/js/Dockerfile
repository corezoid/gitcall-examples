FROM node:22-alpine as builder

WORKDIR /app
COPY src src

RUN addgroup --gid 501 usercode && \
    adduser --disabled-password \
    --gecos "" \
    --shell /usr/sbin/nologin \
    --ingroup usercode \
    --uid 501 \
    usercode
USER usercode

ENTRYPOINT ["node", "/app/src/main.js"]