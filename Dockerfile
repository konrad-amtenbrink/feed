FROM golang:1.19 as builder
WORKDIR /go/src/
COPY . .
RUN go build -v -o /app .

FROM gcr.io/distroless/base
COPY --from=builder /app /app
COPY --from=builder /go/src/migrations /migrations

EXPOSE 8080

CMD ["/app", "server", "--port", "8080"]