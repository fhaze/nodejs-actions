FROM golang:latest as builder
WORKDIR /build
COPY . /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sample

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/sample .
EXPOSE 3000
CMD ["/app/sample"]