FROM golang:1.18-alpine

RUN mkdir /app

RUN addgroup -S appgroup
RUN adduser -S -D -h /app appuser appgroup
RUN chown -R appuser:appgroup /app
USER appuser

COPY main.go /app

EXPOSE 8080
CMD go run /app/main.go
