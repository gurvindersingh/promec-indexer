[![Build Status](https://travis-ci.org/gurvindersingh/promec-indexer.png)](https://travis-ci.org/gurvindersingh/promec-indexer)

# promec-indexer

`Promec-Indexer` flattens the Pep XML data format from [Comet](http://comet-ms.sourceforge.net/) and ingest data with correct mapping as applied from `template.go`

You can test it using Docker as

```
docker run -ti -v /path/to/test.pep.xml:/test.pep.xml gurvin/promec-indexer promec-indexer -pepxml /test.pep-xml -host http://elasisearch-host.com:9200
```
