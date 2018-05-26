FROM golang:1.8-alpine as build
RUN apk update && apk upgrade && apk add --no-cache git

WORKDIR /usr/app/src/chat
COPY . .

RUN go get github.com/CzarSimon/diplo/backend/pkg/httputil \
           github.com/CzarSimon/diplo/backend/pkg/id \
           github.com/CzarSimon/diplo/backend/chat/pkg \
           github.com/CzarSimon/go-endpoint \
           github.com/gin-gonic/gin \
           github.com/lib/pq

RUN go build

FROM alpine:3.6 as run
WORKDIR /root/
COPY --from=build /usr/app/src/chat/chat chat
CMD ["./chat"]
