FROM swipl:9.2.4 as builder

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

ENTRYPOINT swipl src/main.pl --fork=false --sighup=quit --port="$GITCALL_PORT"