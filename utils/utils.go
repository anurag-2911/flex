package utils

import (
	"assetmgmt/pkg/model"
	"encoding/json"
	"fmt"
	"os"
)

const defaultBatchSize = 1000

func GetBatchSize(configfile string) (int, error) {
	file, err := os.Open(configfile)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return defaultBatchSize, err
	}
	defer file.Close()

	var config model.Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return defaultBatchSize, err
	}
	batchsize := config.BatchSize
	fmt.Println("Batch size:", config.BatchSize)
	return batchsize, nil
}
