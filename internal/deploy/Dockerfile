# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as builder
LABEL authors="qusysert"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN cd cmd && CGO_ENABLED=0 GOOS=linux go build -o /docker-parsep
EXPOSE 8080

CMD ["/app/docker-parsep"]

FROM alpine
EXPOSE 8080
RUN adduser -D -H -h /app parsep
WORKDIR /app
USER parsep

COPY --chown=parsep --from=builder /docker-parsep /app
CMD ["/app/docker-parsep"]

