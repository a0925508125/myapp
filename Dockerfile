FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build API
RUN go build -o /app/myapp ./main.go


# runtime image
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/myapp /app/myapp

EXPOSE 8080

CMD ["/app/myapp"]
