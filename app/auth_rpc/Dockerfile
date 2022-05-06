# FROM golang:alpine AS builder

# LABEL stage=gobuilder

# ENV CGO_ENABLED 0
# ENV GOOS linux
# ENV GOPROXY https://goproxy.cn,direct

# WORKDIR /build/zero

# ADD go.mod .
# ADD go.sum .
# RUN go mod download
# COPY . .
# COPY app/auth_rpc/etc /app/etc
# RUN go build -ldflags="-s -w" -o /app/auth app/auth_rpc/auth.go


FROM golang:alpine

# RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
# ENV TZ Asia/Shanghai

WORKDIR /app
COPY ./auth_rpc ./auth_rpc
COPY ./etc ./etc
# COPY --from=builder /app/auth /app/auth
# COPY --from=builder /app/etc /app/etc

CMD ["./auth_rpc", "-f", "etc/auth.yaml"]
