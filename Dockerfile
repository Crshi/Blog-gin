FROM scratch

WORKDIR $GOPATH/src/github.com/Crshi/blog-gin
COPY . $GOPATH/src/github.com/Crshi/blog-gin

EXPOSE 8088
ENTRYPOINT ["./blog-gin"]