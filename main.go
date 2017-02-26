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
	timeZone = flag.String("timezone", "Europe/Oslo", "Timezone to be used in parsing the date from Pep XML file")
	bulkSize = flag.Int("bulksize", 500, "Number of request to send in one bulk request")
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
	xmlMap, err := readCometXML(*pepxml)
	if err != nil {
		log.Error("Failed in parsing XML data ", err)
		os.Exit(1)
	}

	// Convert XML map to ELS bulk index format
	err = indexELSData(xmlMap, *host, *index, *dataType, *bulkSize, *timeZone)
	if err != nil {
		log.Error("Failed in ingesting data for file ", *pepxml, err)
		os.Exit(1)
	} else {
		log.Info("Successfully indexed data from ", *pepxml)
	}

}
