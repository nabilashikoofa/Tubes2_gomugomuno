# Use an official Golang runtime as a parent image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build ./src/backend/main.go

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the built Go application
CMD ["./main"]
