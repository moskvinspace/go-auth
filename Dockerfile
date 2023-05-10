FROM golang:1.20

WORKDIR /app

COPY . .

RUN go get && go build -o bin .

CMD ["/app/bin"]