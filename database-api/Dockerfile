FROM golang:1.21.3-bullseye as base

WORKDIR /usr/app


from base as dev

# RUN go install github.com/cosmtrek/air@v1.49.0
RUN go install github.com/beego/bee/v2@v2.1.0

COPY go/home/matisin/Proyectos/pharmkt/microservices/docker-compose.yml