FROM golang:1.13.5-alpine3.11

WORKDIR $GOPATH/src/github.com/FullOfOrange/go-restful
COPY . $GOPATH/src/github.com/FullOfOrange/go-restful
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./go-restful"]