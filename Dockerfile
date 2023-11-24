FROM golang:1.20 AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH arm64

WORKDIR /build
COPY . .

RUN go install go.opentelemetry.io/collector/cmd/builder@v0.89.0
RUN builder --config otelcol-builder.yml
RUN ls /build/dist


FROM alpine:latest

COPY --from=builder /build/dist/otelcol-custom /otelcol-custom
RUN touch /tmp/simple.log

CMD ["/otelcol-custom", "--config", "/etc/otelcol/config.yml"]
