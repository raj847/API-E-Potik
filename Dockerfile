FROM golang:1.17-alpine3.14

WORKDIR /minpro_arya

COPY . .

RUN go mod download


RUN go build -o mainfile

EXPOSE 8080

CMD ["./mainfile"]