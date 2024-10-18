FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags "-w -s" -o app

####################

FROM debian:bullseye-slim AS production

RUN apt-get update

# hadolint ignore=DL3045
COPY --from=builder /app/app .

CMD ["./app"]

####################

FROM golang:1.23 AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest

CMD ["air"]
