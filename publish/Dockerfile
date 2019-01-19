FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /var/www/publishen/publish

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o publish -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /publish

WORKDIR /publish

COPY --from=builder  /var/www/publishen/publish .

ENTRYPOINT ["/publishen/publish"]

CMD ["./publish"]