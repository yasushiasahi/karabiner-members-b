FROM golang:1.10.4-alpine3.8

WORKDIR /
ENV GOPATH /go

RUN apk add --no-cache git

COPY . /go/src/github.com/yasushiasahi/karabiner-members/api
RUN go get github.com/go-sql-driver/mysql
RUN cd /go/src/github.com/yasushiasahi/karabiner-members/api && go build
RUN cd /go/src/github.com/yasushiasahi/karabiner-members/api && cp api /usr/local/bin/

CMD ["api"]
