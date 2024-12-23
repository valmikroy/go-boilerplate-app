FROM golang:alpine as build

# this gives statically linked binary
# doc - https://jvns.ca/blog/2021/11/17/debugging-a-weird--file-not-found--error/
ENV CGO_ENABLED=0 

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN go build -o /boilerplate .
 
FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /boilerplate /bin/boilerplate

WORKDIR /app
EXPOSE 8080
CMD ["/bin/boilerplate", "--dev", "service"]
