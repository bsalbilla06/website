# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY . .

# Install curl and tailwindcss
RUN apk add --no-cache curl \
    && ./build/install_tailwind.sh \
    && tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --minify

RUN go build -o /server cmd/server/main.go

# Run stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /server /server
COPY --from=builder /app/web /web

EXPOSE 8080

ENTRYPOINT ["/server"]
