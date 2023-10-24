FROM golang:1.21

WORKDIR /project

COPY . .

RUN go mod tidy
