# FROM golang:1.17-alpine3.14

# WORKDIR /minpro_arya

# COPY . .

# RUN go mod download


# RUN go build -o mainfile

# EXPOSE 8080

# CMD ["./mainfile"]
#BUILD STAGE
FROM golang:1.17-alpine AS builder
WORKDIR /minpro_arya
COPY . .
RUN go mod download
RUN go build -o main main.go

#RUN STAGE
FROM alpine:3.14 
WORKDIR /minpro_arya
COPY --from=builder /minpro_arya/config/config.json . 
COPY --from=builder /minpro_arya/main .
EXPOSE 8000

CMD ["/minpro_arya/main"]
