FROM golang:alpine
# Set the working directory to /app
WORKDIR /app
# Copy the current directory contents into the container at /app
COPY . /app
# Download all the dependencies
RUN go mod download
# Build the Go application
RUN go build -o app ./main.go
# Expose port 8000 to the outside world
EXPOSE 8000
# Run the binary program produced by `go build`
CMD ["./app"]
