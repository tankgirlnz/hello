FROM golang:1.4.2-onbuild

RUN mkdir -p /app

WORKDIR /app

ADD app/app.go app.go

RUN go build ./app.go

CMD ["./app"]