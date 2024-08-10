FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build

EXPOSE 8000

CMD ["./budbot"]