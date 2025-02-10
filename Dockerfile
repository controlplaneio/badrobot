FROM golang:1.23.6-alpine AS builder

RUN echo "badrobot:x:25000:25000:badrobot:/home/badrobot:/sbin/nologin" > /passwd && \
    echo "badrobot:x:25000:" > /group
WORKDIR /badrobot
RUN apk add --no-cache ca-certificates
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -X github.com/controlplaneio/badrobot/cmd.version=$VERSION -X github.com/controlplaneio/badrobot/cmd.commit=$COMMIT" \
    -o badrobot .

# ===

FROM scratch
WORKDIR /home/badrobot
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /passwd /group /etc/
COPY --from=builder /badrobot/badrobot /bin/badrobot
USER badrobot

ENTRYPOINT ["/bin/badrobot"]
