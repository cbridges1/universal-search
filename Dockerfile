FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY cmd ./cmd
COPY internal ./internal
RUN go mod download

COPY *.go ./

RUN go build -o /universal-search ./cmd/app/.

EXPOSE 8080

CMD [ "/universal-search" ]