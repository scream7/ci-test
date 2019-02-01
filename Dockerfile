FROM golang:latest
MAINTAINER minjieh <njhmj820@gmail.com>

COPY main.go /app/main.go
WORKDIR /app

EXPOSE 8080

CMD ["go","run","/app/main.go"]