# Build bubble in a stock Go builder container
FROM golang:1.16-alpine3.13 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers g++ llvm bash cmake git gmp-dev openssl-dev

ADD . /bubble
RUN cd /bubble && make clean && make bubble

# Pull bubble into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates libstdc++ bash tzdata gmp-dev
COPY --from=builder /bubble/build/bin/bubble /usr/local/bin/

VOLUME /data/bubble
EXPOSE 6060 6789 6790 6791 16789 16789/udp
CMD ["bubble"]