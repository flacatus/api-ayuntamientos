FROM golang:1.12.1 AS build
WORKDIR /golang
COPY . .

RUN go get github.com/gorilla/mux \
           github.com/google/uuid \
           github.com/gorilla/handlers \
           gopkg.in/mgo.v2 \
           gopkg.in/mgo.v2/bson \
           github.com/sirupsen/logrus

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:3.8
ENV PORT=6000
COPY --from=build /golang/main .
ENTRYPOINT [ "./main" ]