FROM golang:1.10-alpine as build
RUN apk update && apk upgrade && apk add git curl

# Copy source
WORKDIR /go/src/directory
COPY . .

# Install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install dependencies
RUN dep ensure

# Build application
WORKDIR /go/src/directory/server
RUN go build

FROM alpine:3.6 as run
WORKDIR /root/
COPY --from=build /go/src/directory/server/server server
CMD ["./server"]
