FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go build -mod=vendor -v -o app ./cmd
RUN mv app /usr/local/bin

CMD ["app"]