FROM golang:1.16
ENV VERSION=testEnvThisFromEnvVersion1.0
LABEL multi.label1="toygoweb" multi.label2="homework" other="yangqi"
WORKDIR $GOPATH/src/github.com/muyifang1/golang/homework/httpserver/main
ADD . $GOPATH/src/github.com/muyifang1/golang/homework/httpserver
RUN go build .
EXPOSE 8088
ENTRYPOINT /httpserver