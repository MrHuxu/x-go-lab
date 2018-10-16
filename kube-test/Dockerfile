FROM golang:latest

EXPOSE 11011

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/MrHuxu/x-go-lab/kube-test
COPY . /go/src/github.com/MrHuxu/x-go-lab/kube-test
RUN dep ensure -v && go build main.go

ENTRYPOINT [ "./main" ]