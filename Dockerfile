FROM golang:1.7
RUN mkdir $HOME/opt
ENV GOPATH=$HOME/opt
RUN go get github.com/tahmmee/cbdozer
WORKDIR $GOPATH/src/github.com/tahmmee/cbdozer
RUN go build
ENTRYPOINT ["./cbdozer"]
