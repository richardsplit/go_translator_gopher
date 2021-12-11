FROM golang:1.16-alpine

RUN mkdir /go_translator_gopher
ADD . /go_translator_gopher
WORKDIR /go_translator_gopher
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /go_translator_gopher
CMD ["/go_translator_gopher"]