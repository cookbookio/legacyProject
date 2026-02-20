# 1. Builder Stage
FROM golang:1.25-alpine AS builder

# We need build-base (which includes gcc and musl-dev) to compile CGO for SQLite
RUN apk add --no-cache build-base

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# DOWNLOAD & TIDY: This is the critical fix. 
# It ensures all modules in your code are added to the mod file.
RUN go mod download
RUN go mod tidy 

# Copy the rest of the source code
COPY . .

# Build the app with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/cookbook

# 2. Final Run Stage
FROM alpine:latest

# SQLite doesn't usually need the full sqlite package installed to run the binary, 
# but it DOES need the standard C libraries that Go linked against.
RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/cookbook /app/cookbook

# Copy static assets and database
COPY static/ static/
COPY templates/ templates/
COPY demo.db .
COPY api-schema.yaml .

EXPOSE 3000

CMD ["./cookbook"]