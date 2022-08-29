FROM golang

WORKDIR /app

ENV GOPATH=/

ENV GO111MODULE="on"

ENV DB_PASSWORD="qwerty"

COPY . .

# psql
RUN apt-get update

RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-db.sh

# packages
RUN go mod download

RUN go get -u -d github.com/golang-migrate/migrate


RUN go build -o man-utd ./cmd/main/app.go

CMD [  ]


