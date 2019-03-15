FROM arm64v8/golang:1.12

WORKDIR /go/src/app
COPY . .

RUN apt-get install -y \
    libsane-dev \
    imagemagick-dev

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .

EXPOSE 8000

CMD ["./app"]
