FROM golang:alpine AS builder

ARG go_token

RUN apk --no-cache add git openssl
RUN git config --global url."https://oauth2:$go_token@github.com/".insteadOf "https://github.com/"

ENV GOPRIVATE="github.com/cocov-ci"
ENV CGO_ENABLED=0

RUN mkdir /app
WORKDIR /app
COPY golangci-lint golangci-lint
COPY common common
WORKDIR /app/golangci-lint
RUN go build -o ../main ./cmd/main.go

FROM golang:alpine

COPY --from=builder /app/main /bin/plugin-golangci-lint

RUN apk --no-cache add git gcc musl-dev
RUN apk --no-cache add --virtual deps curl

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh

RUN apk del deps

RUN mv bin/golangci-lint /bin/golangci-lint

RUN addgroup -g 1000 cocov && \
    adduser --shell /bin/ash --disabled-password \
   --uid 1000 --ingroup cocov cocov

USER cocov

CMD ["plugin-golangci-lint"]

