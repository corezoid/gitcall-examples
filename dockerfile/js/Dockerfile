FROM node:22-alpine as builder

WORKDIR /app
COPY package.json /app/package.json
COPY package-lock.* /app/
RUN npm i --production
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