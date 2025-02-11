FROM golang:1.22

ENV redisHost=redis

RUN mkdir /app
COPY src/go.mod /app/
COPY src/*.go /app/

WORKDIR /app
RUN go get server
RUN go build -o main .
CMD ["/app/main"]
