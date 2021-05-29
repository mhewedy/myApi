#https://www.callicoder.com/docker-golang-image-container-example/
FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go get -u github.com/swaggo/swag/cmd/swag && swag init -g cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/api/main.go


FROM alpine
WORKDIR /app
COPY --from=builder /app/api /app/app.conf /app/db/migrations ./
ENTRYPOINT ["/app/api"]
