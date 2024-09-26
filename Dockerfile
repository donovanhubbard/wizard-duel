FROM golang:1.23.1-bookworm as builder

RUN mkdir /app
WORKDIR /app

COPY . /app

RUN go mod download
RUN go build -o wizard-duel .

FROM debian:bullseye-slim

RUN useradd -ms /bin/bash newuser

RUN mkdir /app
WORKDIR /app

COPY --chown=app:app --from=builder /app/wizard-duel /app/wizard-duel

USER app

CMD ["/app/wizard-duel"]
