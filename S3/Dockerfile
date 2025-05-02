FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o s3-app

FROM alpine
WORKDIR /app
COPY --from=builder /app/s3-app .

# Create data directory with correct permissions
RUN mkdir -p /data && \
    chown -R 1000:1000 /data && \
    chmod -R 750 /data

USER 1000
EXPOSE 9000
CMD ["./s3-app"]