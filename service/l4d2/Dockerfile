FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/l4d2/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/l4d2 service/l4d2/l4d2.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/l4d2 /app/l4d2
COPY --from=builder /app/etc /app/etc

EXPOSE 8888

CMD ["./l4d2", "-f", "etc/l4d2.yaml"]
