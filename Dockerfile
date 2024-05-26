# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code.
COPY *.go ./

# Copy the templates directory
COPY templates/ ./templates/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-main

# the application is going to listen on by default.
EXPOSE 8080

# Run
CMD ["/docker-main"]
