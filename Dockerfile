FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/Crshi/Blog
COPY . $GOPATH/src/github.com/Crshi/Blog
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./blog"]