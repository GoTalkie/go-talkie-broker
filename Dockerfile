FROM golang:1.20 AS builder
WORKDIR /go/src/github.com/GoTalkie/go-talkie-broker/
COPY ./ ./
RUN go mod vendor
RUN env GOOS=linux CGO_ENABLED=0 go build -buildvcs=false -o brokerApp ./cmd/broker/

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /go/src/github.com/GoTalkie/go-talkie-broker/brokerApp /app
CMD ["/app/brokerApp"]