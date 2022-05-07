FROM golang:1.16

ENV GOPATH=

WORKDIR /api

COPY go.mod .
COPY go.sum .

RUN go mod download

WORKDIR /api/app

COPY . /api/app/

CMD ["go", "run", "."]

EXPOSE 8080