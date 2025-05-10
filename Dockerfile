FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod /.

RUN go mod download
COPY . .
RUN go build -o go-calculator-cmd ./cmd/calculator

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./app/go-calculator-cmd .
ENTRYPOINT ["./go-calculator-cmd"]