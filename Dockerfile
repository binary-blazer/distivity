# Use golang:1.23.3-alpine as the base image
FROM golang:1.23.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source files and go.mod to the container
COPY . .

# Run go build to compile the application
RUN go build -o distivity

# Set the entry point to run the compiled application
ENTRYPOINT ["./distivity"]
