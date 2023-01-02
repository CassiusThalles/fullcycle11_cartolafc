FROM golang:latest

WORKDIR /go/home

COPY . /go/home/

RUN apt-get update
RUN apt-get install -y librdkafka-dev
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

RUN go mod download

EXPOSE 8080

CMD [ "./.docker/start.dev.sh" ]
# CMD ["tail", "-f", "/dev/null"]