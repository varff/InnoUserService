# syntax=docker/dockerfile:1

FROM golang:1.17-alpine  AS builder
RUN mkdir /app
RUN mkdir /tmp
WORKDIR /tmp
ADD . .

RUN go mod download

RUN go build -o main
WORKDIR /app
COPY configs configs/
COPY --from=builder /app /tmp
CMD [ "/app/main" ]