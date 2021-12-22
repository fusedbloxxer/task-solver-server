# syntax=docker/dockerfile:1

##
## Build
##
# Inherit from golang image: https://hub.docker.com/_/golang
FROM golang:1.16-alpine

# Set the new folder for subsequent commands
WORKDIR /server

# Copy the files containing the dependencies
COPY go.mod ./
COPY go.sum ./

# Download the dependencies
RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the environment file
COPY ./env/ ./env/

# !WARNING! Copy the firebase credentials !WARNING!
COPY ./secrets/ ./secrets/

# Copy the source code
COPY ./src/ ./src/

# Expose the 8080 port
EXPOSE 8080

# Generate Swagger / openAPI docs
RUN swag init --dir ./src

# Compile the application
RUN go build -o ./bin/server ./src

# Specify the command to be run when starting the image as a container
CMD ["./bin/server"]