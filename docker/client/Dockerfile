FROM golang:1.21.6 as build

WORKDIR /app

COPY go.sum .
COPY go.mod .

RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/client ./cmd/client/

FROM scratch

COPY --from=build /app/bin/client /bin/client

CMD ["./bin/client/client"]
