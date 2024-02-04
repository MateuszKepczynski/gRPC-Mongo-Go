FROM golang:1.21.6 as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server ./cmd/server/

FROM scratch

COPY --from=build /app/bin/server /bin/server

CMD ["./bin/server/server"]
