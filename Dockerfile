FROM golang:1.18.2 AS builder

WORKDIR /badrobot

RUN apt-get update && apt-get install -y jq

COPY . .

RUN make test && \
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o badrobot .

# ===

FROM alpine:3.15.4

RUN addgroup -S badrobot \
    && adduser -S -g badrobot badrobot \
    && apk --no-cache add ca-certificates

WORKDIR /home/badrobot

COPY --from=builder /badrobot/badrobot /bin/badrobot


USER badrobot

ENTRYPOINT ["/bin/badrobot"]
