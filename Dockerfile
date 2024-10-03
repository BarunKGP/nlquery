# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY server/* ./

# Build
RUN make docker-build

EXPOSE 8001

# Run 
CMD ["/nlquery"]

# RUN mkdir /app
#
# ADD . /app
#
#
# RUN go build -o main ./server/main.go
#
# EXPOSE 8001
# CMD [ "/app/main" ]
