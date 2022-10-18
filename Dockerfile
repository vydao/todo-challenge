FROM golang:1.17-alpine3.15 AS builder

RUN mkdir /app
WORKDIR /app
ADD . .

RUN apk update \
    && apk add build-base \
    && apk add --no-cache git \
    && apk add --no-cache ca-certificates \
    && apk add --update gcc musl-dev \
    && update-ca-certificates

RUN go build -o /app/main /app/cmd/main.go


FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY db/migration ./db/migration
COPY app.env .

# ARG REDIS_SOURCE
# ARG DB_SOURCE
# ARG JWT_SECRET

# ENV rediscon=$REDIS_SOURCE
# ENV postgrescon=$DB_SOURCE
# ENV jwtsecret=$JWT_SECRET

EXPOSE 8080

ENTRYPOINT ["/app/main"]