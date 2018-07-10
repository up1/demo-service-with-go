FROM golang:1.10.3-alpine3.7
WORKDIR /go
COPY . /go
RUN go build -o main main
CMD ["/go/main"]
