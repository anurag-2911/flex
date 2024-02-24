package main

import (
	"assetmgmt/pkg/assetprocessor"
	"assetmgmt/pkg/commandline"
	"assetmgmt/utils"
	"log"
)

var batchsize int = 1000 // read from config file

func init() {
	var err error
	batchsize, err = utils.GetBatchSize("../../config/config.json")
	if err != nil {
		log.Printf("error in reading config file: %v", err)
	}
}

func main() {
	run()
}

func run() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	appid, filePath := commandline.GetCommandLineArguments()
	log.Println("application ID and filepath: ", appid, filePath)
	assetprocessor.ProcessAssets(appid, filePath, batchsize)
}
