FROM golang:1.11-alpine3.9 as builder

ENV APP_DIR $GOPATH/src/github.com/robertke/orders-service

WORKDIR ${APP_DIR}

COPY . .

RUN apk add --update git make
RUN make deps & make build

FROM alpine:3.9

COPY --from=builder /go/src/github.com/robertke/orders-service/out/api /

EXPOSE 8080

CMD ["./api"]
