# Start from the latest golang base image
FROM golang:latest AS build-stage
# FROM golang:latest

# Update dependencies
RUN apt-get update
RUN apt-get install -y ca-certificates

# Add Maintainer Info
LABEL maintainer="Jacob D Cruz <carljacobdiazcruz@gmail.com>"

# Add local directory
ADD . /go/src

# Change working dir
WORKDIR /go/src

# Get GO Dependencies
RUN go get -d -v .
RUN go install -v .

# Build app
RUN go build -o main .

# Deployment to ubuntu image
FROM ubuntu:latest

# Update dependencies
RUN apt-get update
RUN apt-get install -y ca-certificates

# Transfer build to ubuntu image
COPY --from=build-stage /go/src/main .
COPY --from=build-stage /go/src/.env .

# RUN endpoint
CMD ["./main"]
