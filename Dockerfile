FROM golang:alpine

WORKDIR /api

COPY . .

RUN go mod download

RUN go build -o api

CMD [ "api" ]