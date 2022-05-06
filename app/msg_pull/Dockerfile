FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY app/msg_pull/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/msg_pull app/msg_pull/msg_pull.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/msg_pull /app/msg_pull
COPY --from=builder /app/etc /app/etc

CMD ["./msg_pull", "-f", "etc/msg_pull.yaml"]
