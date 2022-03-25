FROM golang:1.17

COPY . .
RUN go build -o score cmd/score
ENTRYPOINT ["/score"]