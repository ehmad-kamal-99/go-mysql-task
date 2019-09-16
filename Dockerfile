FROM golang:latest
RUN go get "github.com/go-sql-driver/mysql"
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]