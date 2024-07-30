
FROM amazoncorretto:21-alpine

WORKDIR /app
COPY gradle gradle
COPY settings.gradle settings.gradle
COPY gradlew gradlew
COPY runner runner

RUN ./gradlew build

RUN addgroup --gid 501 usercode && \
    adduser --disabled-password \
    --gecos "" \
    --shell /usr/sbin/nologin \
    --ingroup usercode \
    --uid 501 \
    usercode
USER usercode

ENTRYPOINT ["java", "-jar", "runner/build/libs/runner.jar"]