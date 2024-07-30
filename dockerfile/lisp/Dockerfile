FROM fukamachi/sbcl:2.4.7 as builder

RUN apt update -y && apt install -y libev4 gcc

WORKDIR /app
COPY src src

RUN addgroup --gid 501 usercode && \
    adduser --disabled-password \
    --gecos "" \
    --shell /usr/sbin/nologin \
    --home /usercode \
    --ingroup usercode \
    --no-create-home \
    --uid 501 \
    usercode
RUN mv /root/ /usercode/
RUN chown -R usercode:usercode /usercode && chmod 0750 /usercode
RUN chown -R usercode:usercode /app && chmod 0750 /app

USER usercode

RUN ros install woo jonathan
RUN ros build /app/src/main.ros 

ENTRYPOINT ["/app/src/main"]
