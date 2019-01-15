# Using this for diferenciating builds as a multi-stage build
FROM golang:1.11-alpine AS builder

WORKDIR $GOPATH/src/github.com/gbrls/VideoRooms

COPY . .

RUN GO111MODULE=on GOPROXY=https://goproxy.io CGO_ENABLED=0 go build -o /bin/VideoRooms

FROM alpine

COPY --from=builder /bin/VideoRooms/ /bin/VideoRooms
COPY --from=builder /go/src/github.com/gbrls/VideoRooms/cfg.yaml /config/cfg.yaml
COPY --from=builder /go/src/github.com/gbrls/VideoRooms/html ./html

EXPOSE 8080

CMD ["VideoRooms", "-cfg=/config/cfg.yaml"]