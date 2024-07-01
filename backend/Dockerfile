# Use the official Golang image as base
FROM golang:1.22.1

# Set the working directory inside the container
WORKDIR /backend

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go build -o backend .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./backend"]
