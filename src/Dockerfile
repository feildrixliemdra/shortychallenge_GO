FROM golang:1.11.1-alpine3.8

RUN apk update && \
    apk upgrade && \
    apk add bash git

WORKDIR /app
COPY . ./

RUN go get github.com/markbates/refresh && go get -v -d

EXPOSE 3000
CMD refresh init && refresh run