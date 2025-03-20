FROM golang:1.24 AS builder

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY backend /app

RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy
RUN go build -o /app/exp

FROM debian:latest AS runner

WORKDIR /app

COPY --from=builder /app/backend /app/backend
