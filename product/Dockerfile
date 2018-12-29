FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /var/www/publishen/product

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /publishen

WORKDIR /publishen

COPY --from=builder  /var/www/publishen/product .

CMD ["./product"]



