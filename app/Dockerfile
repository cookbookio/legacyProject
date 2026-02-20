FROM golang:alpine AS builder
# We DON'T need build-base/gcc anymore!
WORKDIR /app
COPY . .
# Disable CGO and build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cookbook

FROM alpine:latest
# We DON'T need libc6-compat anymore either!
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/cookbook /app/cookbook
COPY static/ static/
COPY templates/ templates/
COPY demo.db .
COPY api-schema.yaml .
EXPOSE 3000
CMD ["./cookbook"]