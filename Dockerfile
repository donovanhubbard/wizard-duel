FROM golang:1.23.1-bookworm

RUN mkdir /app
WORKDIR /app

COPY . /app

RUN go mod download
RUN go build -o wizard-duel .

CMD ["/app/wizard-duel"]