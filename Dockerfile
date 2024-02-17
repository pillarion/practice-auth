# syntax=docker/dockerfile:1

FROM golang:1.22.0-alpine3.19 as builder
WORKDIR /opt/app/
RUN apk update && apk upgrade --available && \
    apk add make && \
    apk add --no-cache tzdata && \
    apk --no-cache add ca-certificates && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    "app-user"
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify
COPY . .
RUN go build -o ./bin/auth-service cmd/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
COPY --from=builder /opt/app/bin/auth-service /go/bin/auth-service
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
ENV TZ Europe/Moscow
USER app-user:app-user
ENTRYPOINT ["go/bin/auth-service"]
