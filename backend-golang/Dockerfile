FROM golang

RUN mkdir /app
COPY backend-golang.go /app

WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
