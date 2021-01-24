FROM golang:1.15 as builder
WORKDIR /go/src/teama-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM scratch
COPY --from=builder /go/src/teama-api/teama-api /usr/bin/
CMD ["teama-api"]
