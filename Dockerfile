# Stage 1: Build the Go application
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

# Stage 2: Run the application with migrations and dependencies
FROM postgres:15 AS migrator
COPY --from=builder /app/migrations /docker-entrypoint-initdb.d/

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ffmpeg
RUN apt-get update && apt-get install -y postgresql-client

WORKDIR /app
COPY --from=builder /app/app .
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]

COPY ./kafka/create-topics.sh /usr/bin/create-topics.sh
RUN chmod +x /usr/bin/create-topics.sh
