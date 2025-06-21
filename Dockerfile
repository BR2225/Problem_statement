# -------- Stage 1: Build --------
FROM golang:1.24.4-alpine AS builder

RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy go source code
COPY . .

# Build the Go binary
RUN go build -o datetime-app

# -------- Stage 2: Run --------
FROM alpine:latest

# Install necessary certificates for HTTPS support (if needed)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/datetime-app .

# Expose the port your app runs on
EXPOSE 8080

# Run the binary
CMD ["./datetime-app"]
