FROM clojure:temurin-22-lein-bookworm-slim

WORKDIR /app

COPY src src
COPY project.clj project.clj

RUN lein install
RUN lein uberjar

# Optimize Java applications launch time
# https://cloud.google.com/run/docs/tips/java
# Run the application, but with a custom trigger that exits immediately.
# In this particular example, the application looks for the '--appcds' flag.
# You can implement a similar flag in your own application.
RUN java -XX:DumpLoadedClassList=/app/classes.lst -jar /app/target/gitcall-1.0.0-standalone.jar --appcds=true; exit 0
# From the captured list of classes (based on execution coverage),
# generate the AppCDS archive file.
RUN java -Xshare:dump -XX:SharedClassListFile=/app/classes.lst -XX:SharedArchiveFile=/app/appcds.jsa --class-path /app/target/gitcall-1.0.0-standalone.jar

RUN addgroup --gid 501 usercode && \
    adduser --disabled-password \
    --gecos "" \
    --shell /usr/sbin/nologin \
    --ingroup usercode \
    --uid 501 \
    usercode
USER usercode

# Set JVM Options for best launch time
# https://cloud.google.com/run/docs/tips/java
ENTRYPOINT ["java", "-Xshare:on", "-XX:SharedArchiveFile=/app/appcds.jsa", "-XX:+TieredCompilation", "-XX:TieredStopAtLevel=1", "-jar", "/app/target/gitcall-1.0.0-standalone.jar"]
