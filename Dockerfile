FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod download

COPY ./ ./

RUN swag init -g ./cmd/backend/main.go -o ./docs
RUN go build ./cmd/backend/main.go

FROM gcr.io/distroless/static-debian12 AS release-stage
WORKDIR /
COPY --from=builder /app/main ./main
USER nonroot:nonroot  
ENTRYPOINT ["/main"]
