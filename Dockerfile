FROM golang:1.23.1-bookworm AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod /app
RUN go mod download

COPY . /app
RUN go build -o wizard-duel .


FROM debian:bookworm-slim


RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/wizard-duel /app/wizard-duel

CMD ["/app/wizard-duel"]
