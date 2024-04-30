# syntax=docker/dockerfile:1

FROM golang:latest

# Set destination for COPY
WORKDIR /home/GolandProjects/todo-app

#RUN apk --no-cache git

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
# COPY ["./todo-app/go.mod", "./todo-app/go.sum", "./"]
COPY ./ ./
RUN go build -o main .
CMD ["./main"]
# RUN go mod download
