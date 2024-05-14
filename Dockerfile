FROM golang:1.22.3 as builder
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

COPY .env .

EXPOSE 3000

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o go-crud main.go

FROM alpine:latest

RUN apk update
RUN apk add --no-cache bash curl jq
WORKDIR /
COPY --from=builder /workspace/go-crud .
COPY --from=builder /workspace/.env .
ENTRYPOINT ["/go-crud"]
