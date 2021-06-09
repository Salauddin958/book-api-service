FROM alpine

COPY dist/book-api-service /bin/

EXPOSE 5001

ENV GOARCH="amd64"
ENV GOOS="linux"
ENV CGO_ENABLED=0

ENTRYPOINT [ "/bin/book-api-service" ]
