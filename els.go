package main

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"github.com/parnurzeal/gorequest"
	"gopkg.in/olivere/elastic.v5"
)

type SearchScore struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SearchHit struct {
	Hit_rank              string        `json:"hit_rank"`
	Peptide               string        `json:"peptide"`
	Peptide_prev_aa       string        `json:"peptide_prev_aa"`
	Peptide_next_aa       string        `json:"peptide_next_aa"`
	Protein               string        `json:"protein"`
	Num_tot_proteins      string        `json:"num_tot_proteins"`
	Num_matched_ions      string        `json:"num_matched_ions"`
	Tot_num_ions          string        `json:"tot_num_ions"`
	Calc_neutral_pep_mass string        `json:"calc_neutral_pep_mass"`
	Massdiff              string        `json:"massdiff"`
	Num_tol_term          string        `json:"num_tol_term"`
	Num_missed_cleavages  string        `json:"num_missed_cleavages"`
	Num_matched_peptides  string        `json:"num_matched_peptides"`
	Search_score          []interface{} `json:"search_score"`
}

type SpectrumQuery struct {
	Spectrum               string `json:"spectrum"`
	Date                   string `json:"date"`
	File                   string `json:"file"`
	FileType               string `json:"filetype"`
	DbName                 string `json:"dbname"`
	SearchEng              string `json:"searcheng"`
	SpectrumNativeID       string `json:"spectrumNativeID"`
	Start_scan             string `json:"start_scan"`
	End_scan               string `json:"end_scan"`
	Precursor_neutral_mass string `json:"precursor_neutral_mass"`
	Assumed_charge         string `json:"assumed_charge"`
	Index                  string `json:"index"`
	Retention_time_sec     string `json:"retention_time_sec"`
	DateTime               time.Time
	Search_result          map[string]interface{} `json:"search_result"`
}

const layout = "2006-01-02T15:04:05"

func indexELSData(
	xmlMap []map[string]interface{},
	host string,
	index string,
	dataType string,
	bulkSize int,
	tz string,
	pepFileName string,
	client *elastic.Client) error {
	loc, _ := time.LoadLocation(tz)
	request := gorequest.New()
	// Create a context
	ctx := context.Background()
	bulkService := elastic.NewBulkService(client)

	// Check is the index exist, then we have template already. Otherwise insert template
	resp, _, _ := request.Get(host + "/" + index).End()
	if resp.StatusCode == 404 {
		resp, _, errs := request.Put(host + "/" + index).
			Send(template).
			End()

		if errs != nil || resp.StatusCode < 200 || resp.StatusCode > 201 {
			errorString := fmt.Sprintf("Could not insert template, Response: %d. Errors: %+v", resp.StatusCode, errs)
			return errors.New(errorString)
		}
	}

	log.Info("Indexing " + pepFileName + " to elasticserch: " + host + " index: " + index)

	for _, specData := range xmlMap {
		var specQuery SpectrumQuery
		var searchHits []SearchHit

		err := mapstructure.Decode(specData, &specQuery)
		if err != nil {
			log.Warn("Failed in parsing SpectrumQuery ", err)
			continue
		}
		err = mapstructure.Decode(specQuery.Search_result["search_hit"], &searchHits)
		if err != nil {
			log.Warn("Failed in parsing Search Result for Index ", specQuery.Index, err)
			continue
		}

		// Convert string date to time and add rank part as seconds to given time
		specQuery.DateTime, err = time.ParseInLocation(layout, specQuery.Date, loc)
		secInt, _ := strconv.Atoi(specQuery.Index)
		sec := time.Duration(secInt) * time.Second
		specQuery.DateTime = specQuery.DateTime.Add(sec)
		if err != nil {
			log.Warn("Failed in parsing time ", err)
			continue
		}

		// Add current spectrum query data to elasticserch data in bulk format
		id := fmt.Sprintf("%x", sha1.Sum([]byte(specQuery.Spectrum)))
		// Add 000 to make time in milliseconds
		jsonData := "{\"@timestamp\":\"" + strconv.FormatInt(specQuery.DateTime.Unix(), 10) + "000" +
			"\", \"database\": \"" + specQuery.DbName +
			"\", \"search_engine\": \"" + specQuery.SearchEng +
			"\", \"file\": \"" + specQuery.File +
			"\", \"filetype\": \"" + specQuery.FileType +
			"\", \"pep_file_name\": \"" + pepFileName +
			"\", \"spectrum\": \"" + specQuery.Spectrum +
			"\", \"spectrumNativeID\": \"" + specQuery.SpectrumNativeID +
			"\", \"start_scan\": \"" + specQuery.Start_scan +
			"\", \"end_scan\": \"" + specQuery.End_scan +
			"\", \"precursor_neutral_mass\": \"" + specQuery.Precursor_neutral_mass +
			"\", \"assumed_charge\": \"" + specQuery.Assumed_charge +
			"\", \"index\": \"" + specQuery.Index +
			"\", \"retention_time_sec\": \"" + specQuery.Retention_time_sec + "\""

		// We have multiple hits, so add them as array with a counter
		cnt := 0
		for _, hit := range searchHits {
			cnt = cnt + 1
			cntS := strconv.Itoa(cnt)
			jsonData = jsonData + ", \"peptide_hit_" + cntS + "\": \"" + hit.Peptide +
				"\", \"peptide_prev_aa_hit_" + cntS + "\": \"" + hit.Peptide_prev_aa +
				"\", \"peptide_next_aa_hit_" + cntS + "\": \"" + hit.Peptide_next_aa +
				"\", \"protein_hit_" + cntS + "\": \"" + hit.Protein +
				"\", \"num_tot_proteins_hit_" + cntS + "\": \"" + hit.Num_tot_proteins +
				"\", \"num_matched_ions_hit_" + cntS + "\": \"" + hit.Num_matched_ions +
				"\", \"tot_num_ions_hit_" + cntS + "\": \"" + hit.Tot_num_ions +
				"\", \"calc_neutral_pep_mass_hit_" + cntS + "\": \"" + hit.Calc_neutral_pep_mass +
				"\", \"massdiff_hit_" + cntS + "\": \"" + hit.Massdiff +
				"\", \"num_tol_term_hit_" + cntS + "\": \"" + hit.Num_tol_term +
				"\", \"num_missed_cleavages_hit_" + cntS + "\": \"" + hit.Num_missed_cleavages +
				"\", \"num_matched_peptides_hit_" + cntS + "\": \"" + hit.Num_matched_peptides + "\""

			var searchScores []SearchScore
			err := mapstructure.Decode(hit.Search_score, &searchScores)
			if err != nil {
				log.Error("Failed in parsing Search Scores ", err)
				continue
			}
			// lets flatten the scores, as we don't want nested arrays
			for _, score := range searchScores {
				switch score.Name {
				case "xcorr":
					jsonData = jsonData + ", \"xcorr_hit_" + cntS + "\":\"" + score.Value + "\""
				case "deltacn":
					jsonData = jsonData + ", \"deltacn_hit_" + cntS + "\":\"" + score.Value + "\""
				case "deltacnstar":
					jsonData = jsonData + ", \"deltacnstar_hit_" + cntS + "\":\"" + score.Value + "\""
				case "spscore":
					jsonData = jsonData + ", \"spscore_hit_" + cntS + "\":\"" + score.Value + "\""
				case "sprank":
					jsonData = jsonData + ", \"sprank_hit_" + cntS + "\":\"" + score.Value + "\""
				case "expect":
					jsonData = jsonData + ", \"expect_hit_" + cntS + "\":\"" + score.Value + "\""
				}
			}
		}

		jsonData = jsonData + "}"
		bulkService.Add(elastic.NewBulkIndexRequest().Index(index).Type(dataType).Id(id).Doc(jsonData))
		if bulkService.NumberOfActions() >= bulkSize {
			_, err := bulkService.Do(ctx)
			if err != nil {
				log.Error("Failed in doing bulk request ", err)
			}
		}
	}
	// ingest the rest of the data
	_, err := bulkService.Do(ctx)
	if err != nil {
		log.Error("Failed in doing bulk request ", err)
	}
	return nil
}

func isFileIndexed(fName string, client *elastic.Client, index string) bool {
	return false
}
