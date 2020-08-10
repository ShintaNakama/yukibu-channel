# syntax = docker/dockerfile:experimental

FROM golang:1.14.7 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
  go build -o ./beetle ./backend/beetle


FROM alpine:latest AS tool

ENV GRPC_HEALTH_PROBE_VERSION=v0.3.2
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe


FROM gcr.io/distroless/static:latest

COPY --from=builder /app/beetle /app/beetle
COPY --from=tool /bin/grpc_health_probe /bin/grpc_health_probe

CMD ["/app/beetle"]
