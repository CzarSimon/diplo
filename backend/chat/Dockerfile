FROM golang:1.10-alpine as build
RUN apk update && apk upgrade && apk add git

# Copy source
WORKDIR /go/src/chat
COPY . .

# Install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install dependencies
RUN dep ensure

# Build application
WORKDIR /go/src/chat/server
RUN go build

FROM alpine:3.6 as run
WORKDIR /root/
COPY --from=build /go/src/chat/server/server server
CMD ["./server"]
