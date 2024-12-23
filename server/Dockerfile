# Build stage
FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN go build -o server ./cmd/main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/server /app/server
CMD ["/app/server"]