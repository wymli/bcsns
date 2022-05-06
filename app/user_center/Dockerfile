FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
COPY . .
# RUN go mod init main
RUN go mod download
COPY app/user_center/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/user_center app/user_center/user_center.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user_center /app/user_center
COPY --from=builder /app/etc /app/etc

# COPY ./user_center .
# COPY ./etc ./etc

CMD ["./user_center", "-f", "etc/user_center.yaml"]
