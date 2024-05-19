FROM golang:1.22.1 AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Construção estática
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/api ./cmd/

FROM alpine:3.19


RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/bin/api .

EXPOSE 8080

CMD ["./api"]
