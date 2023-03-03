FROM golang:1.19 as builder

WORKDIR /go/src/
COPY . .

RUN go build -buildvcs=false -o app

EXPOSE 8080

CMD ["./app", "server", "--port", "8080"]