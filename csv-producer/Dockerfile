FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app
COPY people-10000.csv /usr/local/bin/

COPY wait-for-it.sh /usr/wait-for-it.sh
RUN chmod +x /usr/wait-for-it.sh

CMD ["/bin/bash", "-c", "/usr/wait-for-it.sh --timeout=0 rabbitmq:5672 && app"]
