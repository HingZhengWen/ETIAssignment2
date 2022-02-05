# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./web ./web

RUN go build -o /etiasg2

EXPOSE 8060

CMD [ "/etiasg2" ]