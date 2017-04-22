package main

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
	colorable "github.com/mattn/go-colorable"
	elastic "gopkg.in/olivere/elastic.v5"
)

const srcExtension = "pep.xml"

var (
	pepxml        = flag.String("pepxml", "test.pep.xml", "path to the pepxml file to index")
	host          = flag.String("host", "http://localhost:9200", "Elasticsearch host with port and protocol information")
	index         = flag.String("index", "promec", "Index name in elasticsearch where xml data will be indexed")
	dataType      = flag.String("datatype", "search_hit", "Data type to be used under index")
	timeZone      = flag.String("timezone", "Europe/Oslo", "Timezone to be used in parsing the date from Pep XML file")
	bulkSize      = flag.Int("bulksize", 500, "Number of request to send in one bulk request")
	loglevel      = flag.String("loglevel", "info", "Log level used for printing logs")
	dirName       = flag.String("directory", "", "Directory Path to watch for pepxml files and index")
	sleepInterval = flag.Int64("sleep-interval", 10, "Sleep interval in seconds")
	waitMode      = flag.Bool("wait-mode", false, "Indexer will wait for the file to be created")
	logformat     = flag.String("logformat", "text", "Choose Log format: json or text")
)

func main() {
	flag.Parse()

	// Set up correct log level and format
	if *logformat == "text" {
		log.SetFormatter(&log.TextFormatter{ForceColors: true})
		log.SetOutput(colorable.NewColorableStdout())
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}
	lvl, err := log.ParseLevel(*loglevel)
	if err != nil {
		log.WithFields(log.Fields{
			"detail": err,
		}).Warn("Could not parse log level, using default")
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(lvl)
	}

	interval := time.Duration(*sleepInterval)

	// Create elasticsearch client
	client, err := elastic.NewClient(
		elastic.SetURL(*host),
		elastic.SetSniff(false))

	if err != nil {
		log.Fatal("Failed in creating elasticserch client ", err)
	}

	// We are running in single file indexing mode
	if *dirName == "" {
		log.Info("Promec Indexer started to index file ", *pepxml)

		// Read XML data into a Map
		xmlMap, err := readCometXML(*pepxml, *waitMode)
		if err != nil {
			os.Exit(1)
		}

		// Convert XML map to ELS bulk index format
		err = indexELSData(xmlMap, *host, *index, *dataType, *bulkSize, *timeZone, *pepxml, client)
		if err != nil {
			log.Error("Failed in ingesting data for file ", *pepxml, err)
			os.Exit(1)
		} else {
			log.Info("Successfully indexed data from ", *pepxml)
		}
	} else {
		// Now we are watching the directory for new xml files
		log.Info("Promec Indexer is watching directory \"", *dirName, "\"")
		for {
			//Get the files which are not processed yet
			files, err := watchDir(*dirName, client, *index)
			if err != nil {
				log.Error("Error in watching directory ", err)
				// Sleep predfined interval and retry
				time.Sleep(interval * time.Second)
				continue
			}

			log.Debug("Got files to index ", files)
			for _, file := range files {
				fPath := filepath.Join(*dirName, file)
				xmlMap, err := readCometXML(fPath, *waitMode)
				if err != nil {
					log.Error(err)
					continue
				}
				err = indexELSData(xmlMap, *host, *index, *dataType, *bulkSize, *timeZone, file, client)
				if err != nil {
					log.Error("Failed in ingesting data for file ", file, err)
					continue
				} else {
					log.Info("Successfully indexed data from ", file)
				}
			}
			if len(files) == 0 {
				log.Info("Found no files, sleeping..")
			}
			time.Sleep(interval * time.Second)
		}
	}
}
