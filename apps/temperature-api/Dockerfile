FROM golang:1.22-alpine AS builder
WORKDIR /app


# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .
RUN go build -o temperature-api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/temperature-api .
EXPOSE 8081
CMD ["./temperature-api"]