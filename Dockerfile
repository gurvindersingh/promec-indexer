FROM alpine:3.5

COPY ./promec-indexer /bin/promec-indexer
ENTRYPOINT ["/bin/promec-indexer"]
