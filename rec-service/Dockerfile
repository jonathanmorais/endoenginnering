FROM golang:1.18.2 AS builder

WORKDIR $GOPATH/src/rec-service

COPY . ./

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.12.3

COPY --from=builder /go/src/rec-service/main ./

EXPOSE 30000

ENTRYPOINT ["./main"]