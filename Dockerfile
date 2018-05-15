FROM golang:1.8

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./app.go

CMD ["./app"]