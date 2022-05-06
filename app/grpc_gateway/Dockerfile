FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY app/grpc_gateway/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/grpc_gateway app/grpc_gateway/main.go app/grpc_gateway/config.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/grpc_gateway /app/grpc_gateway
COPY --from=builder /app/etc /app/etc

CMD ["./grpc_gateway", "-f", "etc/gateway.yaml"]
