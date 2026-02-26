FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/main /main

EXPOSE 8080

CMD ["/main"]
