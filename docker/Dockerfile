FROM golang:1.17 as builder
RUN apt update && \
    apt install make

ENV HOME /dalc-test
COPY / ${HOME}
WORKDIR ${HOME}

RUN make build

# stage 2
FROM ubuntu

COPY docker/entrypoint.sh /

COPY --from=builder /dalc-test/dalc-test /

ENTRYPOINT ["/entrypoint.sh"]
CMD ["dalc-test"]
