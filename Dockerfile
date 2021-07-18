FROM golang:latest AS builder
WORKDIR /go/src/target
ADD . .
RUN go env -w GOPROXY="https://goproxy.cn" \
  && go mod download \
  && CGO_ENABLED=0 make

FROM alpine:latest
WORKDIR /srv/target
COPY --from=builder /go/src/target/target /go/src/target/target.yaml /go/src/target/hello.json /srv/target/
VOLUME ["/srv/target/"]
EXPOSE 8080
CMD ["./target", "--config", "target.yaml"]