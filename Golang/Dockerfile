FROM golang:latest

WORKDIR /home/

COPY . /home/

RUN apt-get update &&\
    apt-get install -y librdkafka-dev &&\
    go mod download

EXPOSE 8080

CMD [ "go", "run", "cmd/cartola/main.go" ]