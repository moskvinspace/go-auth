FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY /pkg ./pkg

RUN go get
RUN go build -o bin .

CMD ["/app/bin"]