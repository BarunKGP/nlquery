FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./server/main.go

EXPOSE 8001
CMD [ "/app/main" ]
