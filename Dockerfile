FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache make git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build/api
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

COPY .env .

EXPOSE 8088

CMD ["./app"]