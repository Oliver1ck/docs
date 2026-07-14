ARG GO_VER=1.24.5
ARG ALPINE_VER=3.21
FROM golang:${GO_VER}-alpine${ALPINE_VER}

ARG TG_VER=2.3.95
ENV TG_VER=${TG_VER}

WORKDIR /app

RUN apk add --no-cache git ca-certificates

RUN go install golang.org/x/tools/cmd/goimports@v0.35.0 \
    && go install github.com/seniorGolang/tg/v2/cmd/tg@v${TG_VER}

ENTRYPOINT ["tg"]