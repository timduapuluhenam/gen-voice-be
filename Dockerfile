#Builder
FROM golang:1.17-alpine as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod tidy -compat=1.17
RUN go build -o main .

FROM alpine:3.14
WORKDIR /root/
RUN apk add --no-cache tzdata
COPY --from=builder /app/app/configs/config.json .
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]