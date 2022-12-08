ARG ALPINE_VERSION=3.16
FROM golang:1.19.4-alpine${ALPINE_VERSION} as builder
WORKDIR /go/src/app
COPY . /go/src/app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o check cli/check/cmd/main.go && \
    go build -o in cli/in/cmd/main.go && \
    go build -o out cli/out/cmd/main.go

FROM alpine:$ALPINE_VERSION
RUN apk add --update-cache git openssh-client \
 && git --version
COPY --from=builder /go/src/app/check /opt/resource/
COPY --from=builder /go/src/app/in /opt/resource/
COPY --from=builder /go/src/app/out /opt/resource/
