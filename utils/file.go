package utils

import (
	"encoding/json"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func LoadJsonFile(filename string) map[string]string {
	m := make(map[string]string)
	jsonFile, err := os.Open("./jsons/" + filename + ".json")
	if err != nil {
		log.WithField("method", "os.Open").Error(err)
		return m
	}
	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.WithField("method", "io.ReadAll").Error(err)
		return m
	}
	err = json.Unmarshal(jsonData, &m)
	if err != nil {
		log.WithField("method", "Unmarshal").Error(err)
		return m
	}
	return m
}
