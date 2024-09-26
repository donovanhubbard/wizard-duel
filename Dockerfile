FROM golang:1.23.1-bookworm AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod /app
RUN go mod download

COPY . /app
RUN go build -o wizard-duel .


FROM debian:bookworm-slim

RUN useradd -ms /bin/bash app

RUN mkdir /app && mkdir /app/.ssh && chown -R app:app /app/.ssh
WORKDIR /app

COPY --chown=app:app --from=builder /app/wizard-duel /app/wizard-duel

USER app

CMD ["/app/wizard-duel"]
