# Go Image https://hub.docker.com/_/golang 
FROM golang:alpine

WORKDIR /api

COPY api/go.mod api/go.sum ./

COPY .env ./

RUN go mod download

COPY api /api

# https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
# go -> golang command
# build -> to build the app into a binary
# -o -> argument forces to overwrite any already existent binary
# 'main' -> is the package that I'll be build
RUN go build -o main

# Exposing the same port I've configured
# in my nginx.conf file
EXPOSE 80

CMD ["./main"]
