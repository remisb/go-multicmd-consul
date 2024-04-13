FROM golang:1.22-alpine AS builder

WORKDIR /app

# fetch dependencies
COPY ./go.*  ./
RUN GOPROXY=https://proxy.golang.org go mod download

# build
COPY ./ ./
WORKDIR /app/cmd/server
RUN CGO_ENABLED=0 go build -o server .

FROM alpine AS server_runtime
COPY --from=builder /app/cmd/server /app
EXPOSE 8080
ENTRYPOINT ["./app/server"]