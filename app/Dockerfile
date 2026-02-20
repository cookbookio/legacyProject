# ---------------------------------------------------------
# STAGE 1: The Builder
# ---------------------------------------------------------
FROM golang:1.21-alpine AS builder

# build-base contains gcc, make, and musl-dev needed for CGO (SQLite)
RUN apk add --no-cache build-base

WORKDIR /app

# 1. Copy the dependency maps first
COPY go.mod go.sum ./

# 2. Copy the source code BEFORE tidy
# This allows 'go mod tidy' to see your imports in dbutils.go
COPY . .

# 3. Fetch missing modules and tidy up
RUN go get modernc.org/sqlite/v2
RUN go mod tidy
RUN go mod download

# 4. Build the binary
# CGO_ENABLED=1 is required for many SQLite drivers
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/cookbook

# ---------------------------------------------------------
# STAGE 2: The Final Tiny Image
# ---------------------------------------------------------
FROM alpine:latest

# libc6-compat is often needed for binaries compiled with CGO on Alpine
RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/cookbook /app/cookbook

# Copy all your project assets
COPY static/ static/
COPY templates/ templates/
COPY demo.db .
COPY api-schema.yaml .

# Standardize on Port 3000
EXPOSE 3000

# Run the app
CMD ["./cookbook"]