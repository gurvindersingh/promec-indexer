FROM alpine:3.5

COPY ./promec-indexer /promec-indexer
CMD ["./promec-indexer"]