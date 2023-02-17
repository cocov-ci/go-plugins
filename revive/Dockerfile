FROM golang:alpine AS builder

ARG go_token

RUN apk --no-cache add git openssl
RUN git config --global url."https://oauth2:$go_token@github.com/".insteadOf "https://github.com/"

ENV GOPRIVATE="github.com/cocov-ci"
ENV CGO_ENABLED=0

RUN mkdir /app
WORKDIR /app
COPY revive revive
COPY common common
WORKDIR /app/revive
RUN go build -o ../main ./cmd/main.go

FROM golang:alpine

COPY --from=builder /app/main /bin/plugin-revive

RUN apk --no-cache add git gcc musl-dev
RUN apk --no-cache add --virtual deps curl

RUN go install github.com/mgechev/revive@latest

RUN apk del deps

RUN mv bin/revive /bin/revive

RUN addgroup -g 1000 cocov && \
    adduser --shell /bin/ash --disabled-password \
   --uid 1000 --ingroup cocov cocov

RUN chown -R cocov:cocov /go

USER cocov

CMD ["plugin-revive"]