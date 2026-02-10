# vi skal både bruge et styresystem og en builder til at kører go-apps, så først finder vi builderen som skal passe til styresystemet
FROM golang:1.25-alpine AS builder

# Install build dependencies for CGO
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# vi kopier og kører go.mod først, ellers bliver den sur over at dependensy ikke er der
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build app
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/cookbook

# alpine er en minimal linux distro som er så lille som mulig som kan kører vores app
# hvorfor den skal have et styresystem, når py og node ikke skal er jeg ikke sikker på
FROM alpine:latest

# Install required dependencies for SQLite
RUN apk add --no-cache sqlite

# det her er et nyt workdir i alpine styresystemet, så vi skal kopiere de filer vi skal bruge til at kører appen over "igen"
WORKDIR /app

COPY --from=builder /app/cookbook /app/cookbook

# Copy static files
COPY static/ static/
COPY templates/ templates/

# Copy the database file
COPY demo.db .

# Copy openapi schema
COPY api-schema.yaml .

EXPOSE 3000

CMD ["./cookbook"]