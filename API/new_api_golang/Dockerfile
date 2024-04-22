FROM golang:latest as builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go
FROM alpine:3.14
ENV LOG_LEVEL=bug
ENV GRPC_PORT=8080
ENV ENVIRONMENT=ops
ENV CPU=777
WORKDIR /app
COPY --from=builder app/app .

EXPOSE 8070

# Command to run the executable
CMD ["./app"]
