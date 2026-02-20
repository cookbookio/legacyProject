# ---------------------------------------------------------
# STAGE 1: The Builder (Building the executable)
# ---------------------------------------------------------
# Using 'golang:alpine' fetches the latest Go version to meet 1.25.6 requirements
FROM golang:alpine AS builder

# build-base includes gcc and tools needed for CGO (SQLite)
RUN apk add --no-cache build-base

WORKDIR /app

# 1. Copy the dependency maps
COPY go.mod go.sum ./

# 2. Copy source code BEFORE tidy so Go can see your imports
COPY . .

# 3. Resolve dependencies
# We use 'go get' to specifically pull the sqlite module causing the error
RUN go get modernc.org/sqlite/v2
RUN go mod tidy
RUN go mod download

# 4. Build the binary with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/cookbook

# ---------------------------------------------------------
# STAGE 2: The Runner (Creating the tiny production image)
# ---------------------------------------------------------
FROM alpine:latest

# libc6-compat provides the standard C library for the Go binary to link against
RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/cookbook /app/cookbook

# Copy all your project assets/files
COPY static/ static/
COPY templates/ templates/
COPY demo.db .
COPY api-schema.yaml .

# Standardize on Port 3000
EXPOSE 3000

# Run the application
CMD ["./cookbook"]