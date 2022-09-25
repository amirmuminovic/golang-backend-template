FROM golang:1.19.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /todoapp

EXPOSE 8080

CMD ["/todoapp"]
