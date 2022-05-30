FROM golang:1.18-alpine AS builder
ARG VERSION=unknown
ARG COMMIT=unknown

WORKDIR /badrobot

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -X github.com/controlplaneio/badrobot/cmd.version=$VERSION -X github.com/controlplaneio/badrobot/cmd.commit=$COMMIT" \
    -o badrobot .

# ===

FROM alpine:3.16.0

RUN addgroup -S badrobot \
    && adduser -S -g badrobot badrobot \
    && apk --no-cache add ca-certificates

WORKDIR /home/badrobot

COPY --from=builder /badrobot/badrobot /bin/badrobot

USER badrobot

ENTRYPOINT ["/bin/badrobot"]
