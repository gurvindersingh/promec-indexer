[![Build Status](https://travis-ci.org/gurvindersingh/promec-indexer.png)](https://travis-ci.org/gurvindersingh/promec-indexer)

# promec-indexer

`Promec-Indexer` flattens the Pep XML data format from [Comet](http://comet-ms.sourceforge.net/) and ingest data with correct mapping as applied from `template.go`

You can test it using Docker as

```
docker run -ti -v /path/to/test.pep.xml:/test.pep.xml gurvin/promec-indexer promec-indexer -pepxml /test.pep-xml -host http://elasisearch-host.com:9200
```

## Local setup

You can setup the whole stack with docker containers. The repo contains docker-compose file. Make sure you have `docker-engine` installed on your machine. The run the command from repo directory as
```
docker-compose up
```

This will bring up the whole stack: Elasticsearch, Kibana and Promec-Indexer. Now if you put any pepXML file `/tmp/indexer` directory, it will be indexed and available for searching in the Kibana at `http://localhost:5601`. Before accessing the Kibana URL, make sure it is running as shown in the logs from `docker-compose up` command.