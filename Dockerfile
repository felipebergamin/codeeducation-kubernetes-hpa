FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN go build -o main
RUN ls

FROM golang:latest

COPY --from=builder /app/main /entrypoint
EXPOSE 8000

ENTRYPOINT [ "/entrypoint" ]
