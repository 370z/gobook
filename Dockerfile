FROM golang:1.17-alpine3.14 AS build

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -a -installsuffix cgo -o gobook-build-file ./cmd/main.go

WORKDIR /
COPY --from=build /app/gobook-build-file /bin/app

EXPOSE 1323

CMD ["app"]
