FROM golang:1.12-alpine AS build_base

WORKDIR /app/server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./server .

FROM bitnami/kubectl:latest

COPY --from=build_base /app/server /app/server

CMD ["/app/server"]