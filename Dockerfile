FROM arm64v8/golang:1.12.0-stretch

WORKDIR /go/src/app
COPY . .

RUN /bin/sh install libsane-dev
    #imagemagick-dev \
    #--no-install-recommends

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .

EXPOSE 8000

CMD ["./app"]
