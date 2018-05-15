#FROM golang:1.4.2-onbuild
FROM golang:1.8

RUN mkdir -p /app

WORKDIR /app

ADD app.sh app.sh

RUN chmod +x app.sh 

EXPOSE 8046

CMD ["./app.sh"]