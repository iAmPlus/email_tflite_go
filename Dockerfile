FROM golang:1.12 AS builder

WORKDIR /src

COPY ./ ./

RUN LD_LIBRARY_PATH=/src/libs CGO_CFLAGS=-I/src CGO_LDFLAGS=-L"/src/libs -lrt -lm" CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./application .

ENV LD_LIBRARY_PATH /src/libs/:$LD_LIBRARY_PATH

ENTRYPOINT ["./application"]

EXPOSE 8080
