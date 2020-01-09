#This is a sample Image

#FROM ubuntu
FROM golang:1.10
MAINTAINER itsmeadityaagarwal@gmail.com

RUN apt-get update

RUN echo y|apt-get install mysql-server
RUN /etc/init.d/mysql start


WORKDIR /go/src/github.com/itsmeadi/cart
COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep init||true
RUN dep ensure -v

RUN ./files/init.sh

CMD ["./files/run.sh"]

EXPOSE 9090

