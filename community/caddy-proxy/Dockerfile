FROM golang:1.22.1 as builder

ARG TARGETOS
ARG TARGETPLATFORM
ARG TARGETARCH

WORKDIR /workspace
COPY . .
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build

FROM caddy/caddy:2.7.6-alpine

ARG VERSION
RUN test -n "$VERSION" || (echo "VERSION arg is not set" && false)
ENV VERSION $VERSION

COPY --from=builder /workspace/local-caddy-generate /usr/bin/local-caddy-generate
COPY --from=builder /workspace/entrypoint.sh /entrypoint.sh
RUN chmod +x /usr/bin/local-caddy-generate /entrypoint.sh

# Run the binary.
ENTRYPOINT ["/entrypoint.sh"]
CMD ["/usr/bin/caddy", "run", "--config", "/config/caddy/Caddyfile"]
