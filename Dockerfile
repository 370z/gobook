FROM golang:1.17-alpine3.14

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -0 /dist

EXPOSE 1323

CMD ["/dist"]
