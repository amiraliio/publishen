FROM golang as builder

WORKDIR /var/www/publishen/product

COPY . .

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN CGO=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app

WORKDIR /app

COPY --from=builder  /var/www/publishen/product .

CMD ["./product"]



