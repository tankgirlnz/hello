#FROM golang:1.4.2-onbuild
FROM golang:1.8

RUN mkdir -p /hello

WORKDIR /hello

# ADD app.sh app.sh
ADD app.sh app.sh

# RUN go get github.com/gin-gonic/contrib/gzip
# RUN go get github.com/Divoli/go-proc/address-proc/structs

# RUN go build -o app.sh hello.go

RUN chmod +x app.sh 

EXPOSE 8046

CMD ["./app.sh"]
