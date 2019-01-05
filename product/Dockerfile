FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /var/www/publishen/product

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o product -mod=vendor -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /publishen

WORKDIR /publishen

COPY --from=builder  /var/www/publishen/product .

ENTRYPOINT ["/publishen/product"]

CMD ["./product"]