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

ADD src /my_app
WORKDIR /my_app

RUN source build.sh

RUN rm -rf /go/src/github.com/go-openapi/spec
RUN go get -v -d

EXPOSE 3000
CMD ["sh", "-c", "go run main.go serve"]