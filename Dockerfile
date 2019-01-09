FROM golang
# COPY ./app /app
WORKDIR .
EXPOSE 5002
RUN go get -t -d -v ./... && go build -v ./...
