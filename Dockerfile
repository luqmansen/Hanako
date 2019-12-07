FROM golang:alpine as builder

LABEL maintainer="Luqman Setyo Nugroho <luqmansen@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /go/src/github.com/luqmansen/hanako
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /go/src/github.com/luqmansen/hanako/main .
COPY --from=builder /go/src/github.com/luqmansen/hanako/.env .

EXPOSE 8000

CMD ["./main"]
