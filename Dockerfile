FROM golang:1.20-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o bin/user_etl ./cmd/user_etl/user_etl.go

ENTRYPOINT ["./bin/user_etl/user_etl"]