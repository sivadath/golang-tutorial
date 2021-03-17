## We specify the base image we need for our
## go application
FROM golang:1.12.0-alpine3.9
## We create an /app directory within our
## image that will hold our application source
## files
## We need gcc for sqlite3 installation
RUN apk add build-base
RUN mkdir /easyblog
## We copy everything in the root directory
## into our /app directory
ADD . /easyblog
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
## We set GOPATH environment variable, not using any package ##dependency management tools.
ENV GOPATH /easyblog
WORKDIR /easyblog
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main ./src/main/main.go
## Our start command which kicks off
## our newly created binary executable
CMD ["/easyblog/main"]
