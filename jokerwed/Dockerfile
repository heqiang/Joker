FROM golang:1.17


WORKDIR $GOPATH/src/test
COPY .. $GOPATH/src/test
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
RUN  go mod tidy &&go build -o jockerweb

EXPOSE 8081
ENTRYPOINT  ["./jockerweb"]