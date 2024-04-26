
FROM golang:1.16 AS build

WORKDIR /app


COPY . .

RUN go mod download



COPY *.go ./


RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

# FROM alpine:3.14

# WORKDIR /app

# COPY --from=build app/app .

CMD ["./app"]


























