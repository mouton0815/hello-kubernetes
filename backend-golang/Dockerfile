FROM golang

ENV redisHost=redis

RUN mkdir /app
COPY src/*.go /app/

WORKDIR /app
RUN go get github.com/gomodule/redigo/redis
RUN go build -o main .
CMD ["/app/main"]
