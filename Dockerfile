FROM golang:1.22

RUN apt-get update && apt-get install -y nodejs npm && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN cd ./frontend && npm install && npm run build

RUN cd ..

RUN go build

EXPOSE 8000

CMD ["./budbot"]