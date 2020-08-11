FROM golang:1.14.4-alpine AS compiler

WORKDIR /app/src

COPY go.mod go.mod
COPY go.sum go.sum
COPY . .

RUN go build -o main

FROM alpine

COPY --from=compiler /app/src/main /app/main

ENTRYPOINT ["/app/main"]
