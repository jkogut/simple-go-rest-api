FROM golang:1.11-alpine
RUN apk update && apk upgrade && apk add --no-cache bash git

ENV GIT_TERMINAL_PROMPT=1
ENV SOURCES /go/src/simple-go-rest-api
COPY . ${SOURCES}

WORKDIR ${SOURCES}
RUN go get -t -d -v ./... && CGO_ENABLED=0 go build

CMD ${SOURCES}/simple-go-rest-api
EXPOSE 5002