# Build 
FROM golang:1.25.5-alpine AS builder
WORKDIR /app

# Install 
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Generate Swagger docs
RUN swag init 

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

COPY --from=builder /app/docs ./docs

RUN chmod +x main
EXPOSE 8080
CMD ["./main"]