FROM golang:1.17-alpine3.14
RUN apk add build-base
WORKDIR /app

COPY . .
RUN go mod download
RUN go build ./cmd/main.go -a -o /gobook-build-file



EXPOSE 1323

CMD ["/gobook-build-file"]
