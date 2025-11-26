# ---- Build Stage ----
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Download dependencies first (better caching)
COPY go.mod ./
RUN go mod download

# Copy source
COPY . .

# Build static executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# ---- Run Stage ----
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/app .

# Run binary
CMD ["./app"]