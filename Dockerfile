FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN  go build -o main ./main.go

EXPOSE 3333
EXPOSE 8080
EXPOSE 8080

CMD ["/app/main"]