# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /home
COPY main .
COPY templates ./templates
EXPOSE 8080 9001

CMD [ "/bin/sh", "-c", "./main" ]