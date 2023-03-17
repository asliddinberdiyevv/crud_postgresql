# Use an official Golang runtime as a parent image
FROM golang:latest AS builder
# Set the working directory to /app
WORKDIR /app
# Copy the current directory contents into the container at /app
COPY . /app
# Download all the dependencies
RUN go mod download
# Build the Go application
RUN go build -o app .
# Use an official Alpine Linux runtime as a parent image
FROM alpine:latest
# Install the PostgreSQL client
RUN apk add --no-cache postgresql-client
# Copy the built binary from the previous stage to the current stage
COPY --from=builder /app/app /app/app
# Set the working directory to /app
WORKDIR /app
# Expose port 8000 to the outside world
EXPOSE 8000
# Run the binary program produced by `go build`
CMD ["./app"]
