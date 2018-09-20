FROM golang:1.10.4-alpine3.8


ENV GOPATH /go
RUN apk add git curl

WORKDIR /go/src/github.com/yasushiasahi/karabiner-members/api

RUN go get github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get github.com/cespare/reflex
RUN go get github.com/PuerkitoBio/goquery

RUN cp /go/bin/reflex /usr/local/bin/

CMD ["reflex","-r","go$","-s","go","run","main.go"]

EXPOSE 3000
