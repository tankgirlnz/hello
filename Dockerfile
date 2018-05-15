FROM golang:1.4.2-onbuild

RUN mkdir -p /app

WORKDIR /app

ADD app/app app

RUN go build ./app.go

CMD ["./app"]