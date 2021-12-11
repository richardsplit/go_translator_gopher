FROM golang:1.17-alpine3.13 as go-build

# Port can be change to your suiting 
WORKDIR /go_translator_gopher
COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod vendor
RUN go mod download
RUN go build -o main main.go

# Run stage 
FROM alpine:3.13
WORKDIR /go_translator_gopher
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["/go_translator_gopher/main"]