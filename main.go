package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
)

var (
	pepxml   = flag.String("pepxml", "test.pep.xml", "path to the pepxml file to index")
	host     = flag.String("host", "http://localhost:9200", "Elasticsearch host with port and protocol information")
	index    = flag.String("index", "promec", "Index name in elasticsearch where xml data will be indexed")
	dataType = flag.String("datatype", "search_hit", "Data type to be used under index")
)

func init() {
	// Log as JSON to stderr
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
}

func main() {
	flag.Parse()
	log.Info("Promec Indexer started to index file ", *pepxml)

	// Read XML data into a Map
	xmlMap, _ := readCommetXML(*pepxml)

	// Convert XML map to ELS bulk index format
	err := indexELSData(xmlMap, *host, *index, *dataType)
	if err != nil {
		log.Error("Failed in ingesting data for file ", *pepxml, err)
	} else {
		log.Info("Successfully indexed data from ", *pepxml)
	}

}
