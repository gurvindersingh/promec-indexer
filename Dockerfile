FROM alpine:3.5

COPY ./promec-indexer /bin/promec-indexer

ENV UID 999
ENV GID 999

USER $UID:$GID

ENTRYPOINT ["/bin/promec-indexer"]
