ARG GO_VERSION=1.18

## Build container
FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --no-cache git

WORKDIR /src

COPY ./go.mod ./
#RUN go mod download # not needed bc 0 deps

COPY ./ ./
RUN go build -o /score /src/cmd/score

## Final container
FROM scratch AS final

COPY --from=builder /src/score.csv /score.csv
COPY --from=builder /score /score

ENTRYPOINT ["/score"]