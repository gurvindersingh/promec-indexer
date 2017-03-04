package main

import (
	"io/ioutil"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/clbanning/mxj"
)

func readCometXML(xmlfile string) ([]map[string]interface{}, error) {
	xmlData, err := readXMLFile(xmlfile)
	mxj.PrependAttrWithHyphen(false)
	mapVal, err := mxj.NewMapXml(xmlData)
	if err != nil {
		log.Error("Failed in parsing xml contents", err)
		return nil, err
	}

	date, err := mapVal.ValuesForKey("date")
	if err != nil {
		log.Error("Failed in getting date", err)
		return nil, err
	}
	file, err := mapVal.ValuesForKey("base_name")
	if err != nil {
		log.Error("Failed in getting file name", err)
		return nil, err
	}
	fileType, err := mapVal.ValuesForKey("raw_data")
	if err != nil {
		log.Error("Failed in getting fileType name", err)
		return nil, err
	}

	xmlVal, err := mapVal.ValuesForKey("spectrum_query")
	if err != nil {
		log.Error("Failed in getting spectrum_query", err)
		return nil, err
	}

	dbVal, err := mapVal.ValuesForKey("local_path")
	if err != nil {
		log.Error("Failed in getting database_name ", err)
		return nil, err
	}

	searchEngVal, err := mapVal.ValuesForKey("search_engine")
	if err != nil {
		log.Error("Failed in getting search_engine ", err)
		return nil, err
	}

	var xmlMap []map[string]interface{}
	for _, val := range xmlVal {
		valMap := val.(map[string]interface{})
		valMap["dbname"] = dbVal[0]
		valMap["searcheng"] = searchEngVal[0]
		valMap["date"] = date[0]
		valMap["file"] = file[0].(string) + fileType[0].(string)
		valMap["fileType"] = strings.Replace(fileType[0].(string), ".", "", 1)
		xmlMap = append(xmlMap, valMap)
	}

	return xmlMap, nil
}

func readXMLFile(xmlfile string) ([]byte, error) {
	for {
		xmldata, err := ioutil.ReadFile(xmlfile)
		if err != nil && strings.Contains(err.Error(), "no such file") {
			log.Debug(xmlfile, " not created/found, sleeping..")
			time.Sleep(10 * time.Second)
			continue
		}
		if err != nil {
			return nil, err
		}
		return xmldata, nil
	}
}
