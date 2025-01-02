GNU nano 7.2                                                                                                Dockerfile                                                                                                          
# Stage 1: Build the Go application
FROM golang:1.21 AS builder
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application and build
COPY . .
RUN go build -o app .

# Stage 2: Create the runtime image
FROM ubuntu:22.04

# Install ffmpeg and dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    ffmpeg

# Set the GOOGLE_CLOUD_PROJECT environment variable
ENV GOOGLE_CLOUD_PROJECT="learnlingo-445821"

# Copy the built Go binary from the builder stage
COPY --from=builder /app/app /usr/local/bin/app

# Copy the Google Cloud credentials file
COPY key.json /app/key.json

# Set GOOGLE_APPLICATION_CREDENTIALS environment variable
ENV GOOGLE_APPLICATION_CREDENTIALS="/app/key.json"


# Expose port 8080 for the Go application
EXPOSE 8080

# Command to run the Go application
CMD ["app"]























[ Read 38 lines ]
^G Help          ^O Write Out     ^W Where Is      ^K Cut           ^T Execute       ^C Location      M-U Undo         M-A Set Mark     M-] To Bracket   M-Q Previous     ^B Back          ^◂ Prev Word     ^A Home
^X Exit          ^R Read File     ^\ Replace       ^U Paste         ^J Justify       ^/ Go To Line    M-E Redo         M-6 Copy         ^Q Where Was     M-W Next         ^F Forward       ^▸ Next Word     ^E End