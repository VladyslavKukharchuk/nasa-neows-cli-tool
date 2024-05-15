FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o nasa-neows-cli-tool .


FROM alpine:edge

WORKDIR /app

COPY --from=builder /app/nasa-neows-cli-tool .

CMD ["/app/nasa-neows-cli-tool"]