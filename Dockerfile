# Use an official Golang runtime as a parent image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go application
CMD ["./main"]
