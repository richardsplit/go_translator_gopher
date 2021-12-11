FROM golang:1.12.0-alpine3.9

RUN mkdir /go_translator_gopher
ADD . /go_translator_gopher
WORKDIR /go_translator_gopher
RUN go build -o main .
CMD ["/go_translator_gopher/main.go"]