FROM golang:1.25 AS builder

ARG ENTER_PATH
ARG APP_NAME

ADD . /app/
WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

RUN export GO111MODULE=on && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app/${APP_NAME} ./${ENTER_PATH}

# -----------------------------------------------------------

FROM alpine:3.16

ARG APP_NAME
ARG EXPOSE_PORT

# 這是為了讓 CGO_ENABLED=0 編譯的 Go 程式在 Alpine 上成功運行的最後嘗試
# 1. ca-certificates & tzdata: 網路和時間處理
# 2. libc6-compat: 解決 Go 靜態二進位檔在 musl C Library 環境中遇到的兼容性問題
RUN apk --no-cache add ca-certificates tzdata libc6-compat
COPY --from=builder /app/${APP_NAME} /app-exec
RUN chmod +x /app-exec
EXPOSE ${EXPOSE_PORT}
CMD ["/app-exec"]
