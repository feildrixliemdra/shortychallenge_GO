FROM golang:1.11.5-alpine3.9

RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    apk add gcc && \
    apk add musl-dev && \
    apk add curl

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /my_app

ADD src /go/src/my_app
WORKDIR /go/src/my_app

RUN dep ensure -v

RUN cp -rf vendor/* /go/src
RUN rm -rf vendor

EXPOSE 3000
CMD ["sh", "-c", "swag init && go run main.go serve"]