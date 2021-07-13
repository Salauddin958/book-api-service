FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/Salauddin958/book-api-service/
WORKDIR /go/src/github.com/Salauddin958/book-api-service
RUN go mod download
COPY . /go/src/github.com/Salauddin958/book-api-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/book-api-service github.com/Salauddin958/book-api-service

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/Salauddin958/book-api-service /usr/bin/book-api-service
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/book-api-service"]