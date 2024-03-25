FROM golang:latest AS build

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o report_generator ./cmd/report_generator

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/report_generator .

ENTRYPOINT [ "./report_generator" ]