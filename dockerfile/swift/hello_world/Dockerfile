# ================================
# Build stage
# ================================
FROM swift:6.3-noble AS build

ENV DEBIAN_FRONTEND=noninteractive

# System dependencies
RUN apt-get update && apt-get install -y \
    libjemalloc-dev \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /build

# Cache dependencies
COPY Package.* ./
RUN swift package resolve

# Copy source
COPY . .

# Build release
RUN swift build -c release --disable-sandbox

# ================================
# Staging stage
# ================================
WORKDIR /staging

# Copy ONLY executable (not entire .build tree)
RUN BIN_PATH="$(swift build -c release --show-bin-path --package-path /build)" && \
    EXECUTABLE=$(ls "$BIN_PATH" | grep -v -E '\.d|\.swift|\.build' | head -n 1) && \
    cp "$BIN_PATH/$EXECUTABLE" ./App

# Swift backtrace helper
RUN cp /usr/libexec/swift/linux/swift-backtrace-static ./ 2>/dev/null || true

# SPM resources
RUN find "$(swift build -c release --show-bin-path --package-path /build)" \
    -name "*.resources" -exec cp -Ra {} ./ \; || true

# Optional Public/Resources
RUN [ -d /build/Public ] && { mv /build/Public ./Public && chmod -R a-w ./Public; } || true
RUN [ -d /build/Resources ] && { mv /build/Resources ./Resources && chmod -R a-w ./Resources; } || true


# ================================
# Runtime stage
# ================================
FROM debian:trixie-slim

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    ca-certificates \
    tzdata \
    libjemalloc2 \
 && rm -rf /var/lib/apt/lists/*

# Swift runtime
COPY --from=build /usr/lib/swift/linux /usr/lib/swift/linux

RUN useradd --user-group --create-home --system \
    --home-dir /usercode usercode

WORKDIR /app

COPY --from=build /staging /app

ENV SWIFT_BACKTRACE=enable=yes,sanitize=yes,threads=all,images=all,interactive=no,swift-backtrace=./swift-backtrace-static

ENV LOG_LEVEL=info

USER usercode

CMD ["sh", "-c", "/app/App serve --env production --hostname 0.0.0.0 --port ${GITCALL_PORT}"]