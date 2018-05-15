FROM golang:1.4.2-onbuild

RUN mkdir -p /app

WORKDIR /app

ADD app.sh app.sh

RUN chmod +x app.sh 

EXPOSE 8045

CMD ["./app.sh"]