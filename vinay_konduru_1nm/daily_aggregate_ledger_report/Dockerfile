FROM golang:latest AS build

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o daily_aggregate_ledger_report ./cmd/daily_aggregate_ledger_report

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/daily_aggregate_ledger_report .

ENTRYPOINT [ "./daily_aggregate_ledger_report" ]