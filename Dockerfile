FROM golang:1.8	 
 	 
RUN mkdir -p /app	
 	 
WORKDIR /app
 	 
ADD app.sh app.sh



# #FROM golang:1.4.2-onbuild
# FROM golang:1.8

# RUN mkdir -p /hello

# WORKDIR /hello

# # ADD app.sh app.sh
# ADD . .

# RUN go get github.com/gin-gonic/contrib/gzip

# RUN go build -o app.sh hello.go

# RUN chmod +x app.sh 

# EXPOSE 8046

# CMD ["./app.sh"]
