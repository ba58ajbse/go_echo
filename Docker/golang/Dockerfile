FROM golang:1.17.2

RUN apt update

WORKDIR /go/src
ENV GO111MODULE on

EXPOSE 8080

RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls

# echo インストール
# RUN go mod init echo_app
# RUN go get github.com/labstack/echo/v4
