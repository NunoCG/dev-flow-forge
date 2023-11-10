ARG GO_VERSION \
    GO_PORT

FROM golang:${GO_VERSION}

ENV GO_PORT=${GO_PORT}

WORKDIR /devflowapp

COPY . .

RUN go build -o devflowapp

EXPOSE ${GO_PORT}

CMD ["./devflowapp"]