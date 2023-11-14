FROM golang:1.21.3-bullseye 
ARG APP_PORT
ENV APP_PORT ${APP_PORT}
ARG BUILD
ENV BUILD ${BUILD}

WORKDIR /usr/app
COPY go.mod/ ./go.mod
COPY go.sum/ ./go.sum
COPY cmd/ ./cmd
COPY internal/ ./internal
COPY conf/ ./conf
COPY migrations/ ./migrations
COPY scripts/ ./scripts

RUN go mod download

RUN ./run_m



EXPOSE $APP_PORT