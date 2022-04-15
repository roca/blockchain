FROM ubuntu:16.04

WORKDIR /var/project

Copy . .

RUN tar -xvf go1.18.1.linux-amd64.tar.gz
RUN mv go /usr/local/go

ENV GOROOT=/usr/local/go
ENV GOPATH=/var/project/goProject
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

RUN apt-get update
RUN apt-get install -y curl
RUN curl -sl https://deb.nodesource.com/setup_10.x | bash -
RUN apt-get update
RUN apt-get install -y nodejs
RUN apt-get install -y git
RUN apt-get install -y python
RUN apt-get install -y libltdl-dev