FROM golang:1.16 AS build

WORKDIR /app

COPY genproto /usr/local/go/src/genproto

COPY . ./

RUN go mod download

RUN 

COPY *.go ./


RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go



CMD ["./app"]












