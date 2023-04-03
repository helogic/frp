FROM golang:1.20-alpine3.16
WORKDIR /home
COPY ./ ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build ./cmd/frps/
FROM logiche/alpine-tz:3.16
WORKDIR /run
ENV vhost_http_port 0
COPY --from=0 /home/frps /run/app
RUN chmod +x /run/app
CMD ["/run/app"]
