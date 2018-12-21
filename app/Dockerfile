FROM golang
COPY /app /app
WORKDIR /app
EXPOSE 5002
RUN go get -t -d -v ./... && go build -v ./...
