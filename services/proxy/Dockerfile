# ---- Build Stage ----
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o proxy .

# ---- Run Stage ----
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/proxy .

# Expose port
EXPOSE 8080

# Run
CMD ["./proxy"]
