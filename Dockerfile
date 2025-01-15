# Build stage
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Use Distroless for minimal runtime image
FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/server .

EXPOSE 8080

USER nonroot

CMD ["/server"]
