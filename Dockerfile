FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o avito-tech ./cmd/main.go

CMD ["./avito-tech"]