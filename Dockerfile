FROM golang:1.21.5-alpine3.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /app/main main
EXPOSE 8000
CMD [ "./main" ]