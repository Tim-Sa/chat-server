FROM golang:1.21.6-alpine AS builder

COPY . /github.com/Tim-Sa/chat-server/source/
WORKDIR /github.com/Tim-Sa/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/main.go



FROM alpine:3.19.1

WORKDIR /root/
COPY --from=builder /github.com/Tim-Sa/chat-server/source/bin/chat_server .

CMD [ "./chat_server" ]