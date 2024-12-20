FROM golang:alpine AS builder

RUN apk add --no-cache --update git gcc rust

COPY . /src
WORKDIR /src

RUN GOPROXY=direct GOSUMDB=off CGO_ENABLED=0 go build -a -ldflags "-linkmode external -extldflags -static" -o /usr/local/bin/supervisord github.com/couriourc/supervisord-plus

FROM scratch

COPY --from=builder /usr/local/bin/supervisord-plus /usr/local/bin/supervisord-plus

ENTRYPOINT ["/usr/local/bin/supervisord-plus"]
