FROM golang:1.22.12-alpine3.21 AS builder

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /api ./cmd/main.go

FROM scratch

COPY --chown=0:0 --from=builder /api /
EXPOSE 8080

USER 65534

ENTRYPOINT ["/api"]
