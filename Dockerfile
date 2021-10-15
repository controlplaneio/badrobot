FROM debian:buster

RUN \
  DEBIAN_FRONTEND=noninteractive \
    apt update && apt install --assume-yes --no-install-recommends \
      bash \
  \
  && rm -rf /var/lib/apt/lists/*
