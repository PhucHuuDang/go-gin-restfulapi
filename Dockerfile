FROM ubuntu:latest
COPY . .
LABEL authors="harrydang"

ENTRYPOINT ["top", "-b"]