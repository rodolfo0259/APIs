FROM golang:1.21.2-alpine as base

WORKDIR /app
COPY . .
RUN go mod tidy

FROM base as deploy

CMD ["go", "run", "cmd/app/main.go"]
