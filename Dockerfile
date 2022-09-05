# Build Stage
# First pull Golang image
FROM golang:alpine as build-env

ENV APP_NAME chuck-norris-challenge
ENV CMD_PATH /cmd/app/main.go

COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage
FROM alpine

# Set environment variable
ENV APP_NAME chuck-norris-challenge

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# Expose application port
EXPOSE 8081sss

# Start app
CMD ./$APP_NAME