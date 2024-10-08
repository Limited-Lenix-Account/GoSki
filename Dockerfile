# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -o /docker-gs-ping

CMD ["/docker-gs-ping"]