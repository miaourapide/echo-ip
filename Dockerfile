FROM golang:1.20-alpine AS base
WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

FROM base AS test
RUN go test

FROM base AS build
RUN go build -v -o app

FROM build AS run
CMD ["./app"]
