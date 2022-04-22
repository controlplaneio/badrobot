FROM golang:1.18.1 AS builder

WORKDIR /badrobot

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o badrobot .

# ===

FROM alpine:3.15.4

RUN addgroup -S badrobot \
    && adduser -S -g badrobot badrobot \
    && apk --no-cache add ca-certificates

WORKDIR /home/badrobot

COPY --from=builder /badrobot/badrobot /bin/badrobot
COPY --from=stefanprodan/kubernetes-json-schema:latest /schemas/master-standalone /schemas/master-standalone-strict
COPY ./templates/ /templates

RUN chown -R badrobot:badrobot ./ /schemas

USER badrobot

ENTRYPOINT ["badrobot"]
CMD ["http", "8080"]
