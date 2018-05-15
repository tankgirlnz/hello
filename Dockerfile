FROM golang:1.4.2-onbuild

RUN mkdir -p /app

WORKDIR /app

ADD app/app.sh app.sh

RUN chmod +x app.sh; 

CMD ["./app.sh"]