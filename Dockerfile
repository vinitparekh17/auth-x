FROM golang:1.21.5-alpine3.18
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GCO_ENABLED=0 GOOS=linux go build -o main .
EXPOSE 8000
CMD ["./main"]