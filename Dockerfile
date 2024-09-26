FROM golang:1.23.1

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# build go app
RUN go mod download
RUN go build -o auth_test ./cmd/main.go

EXPOSE 8088

CMD ["./auth_test"]