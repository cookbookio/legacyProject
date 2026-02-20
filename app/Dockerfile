# ---------------------------------------------------------
# STAGE 1: The Builder
# ---------------------------------------------------------
FROM golang:alpine AS builder

# build-base is required for CGO (SQLite)
RUN apk add --no-cache build-base

WORKDIR /app

# 1. Copy everything (including go.mod AND your .go files)
# We need the .go files so 'go mod tidy' can see what you are actually importing
COPY . .

# 2. Let Go find the right dependencies automatically
# This will fix the 'v2' vs 'v1' issue by looking at your actual code
RUN go mod tidy
RUN go mod download

# 3. Build the binary
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/cookbook

# ---------------------------------------------------------
# STAGE 2: The Runner (The rest remains exactly the same)
# ---------------------------------------------------------
FROM alpine:latest
RUN apk add --no-cache ca-certificates libc6-compat
WORKDIR /app
COPY --from=builder /app/cookbook /app/cookbook
COPY static/ static/
COPY templates/ templates/
COPY demo.db .
COPY api-schema.yaml .
EXPOSE 3000
CMD ["./cookbook"]