FROM alpine:3.5

ENV COMET_VERSION 2016012

RUN apk update && apk add --no-cache ca-certificates wget tzdata
WORKDIR /tmp
RUN ZIP=comet_binaries_$COMET_VERSION.zip && \
	wget -q https://github.com/BioDocker/software-archive/releases/download/Comet/$ZIP -O /tmp/$ZIP && \
	unzip /tmp/$ZIP && \
	chmod +x /tmp/comet_binaries_$COMET_VERSION/comet.$COMET_VERSION.linux.exe && \
	mv /tmp/comet_binaries_$COMET_VERSION/comet.$COMET_VERSION.linux.exe /bin/comet && \
	rm -rf /tmp/comet_binaries*

COPY dist/promec-indexer-linux-amd64 /bin/promec-indexer

ENV UID 999
ENV GID 999

RUN mkdir -p /data
ADD comet-start.sh /bin/
USER $UID:$GID

CMD ["/bin/promec-indexer"]
