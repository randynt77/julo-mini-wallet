FROM golang:alpine as builder


RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o julo-mini-wallet .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/julo-mini-wallet .
COPY --from=builder /app/.env .       

EXPOSE 8080

CMD ["./julo-mini-wallet"]