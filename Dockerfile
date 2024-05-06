FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY . ./

RUN go build -o /nasa-neows-cli-tool ./main.go

CMD ["/nasa-neows-cli-tool"]