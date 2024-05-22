ARG GO_VERSION=1.19-alpine3.15

FROM golang:${GO_VERSION}

ENV GO111MODULE=on

RUN mkdir -p /app

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8080

ENTRYPOINT ["/app/exoplanet"]