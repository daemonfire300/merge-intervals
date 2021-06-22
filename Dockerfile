FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-w -extldflags "-static"' \ 
    -v -o /usr/local/bin/mergectl ./cmd/mergectl

RUN rm -rf /go/src/app

ENTRYPOINT ["mergectl"]