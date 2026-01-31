FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:3.19

WORKDIR /app

COPY --from=build /app/server /app/server
COPY templates /app/templates
COPY db /app/db
COPY styles.css /app/styles.css
COPY assets /app/assets

ENV PORT=8080
EXPOSE 8080

CMD ["./server"]
