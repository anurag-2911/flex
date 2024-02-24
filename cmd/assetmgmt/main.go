package main

import (
	"assetmgmt/pkg/assetprocessor"
	"assetmgmt/pkg/commandline"
	"assetmgmt/utils"
	"log"
)

func main() {
	appid, filePath := commandline.GetCommandLineArguments()
	log.Println("application ID and filepath : ", appid, filePath)
	assetprocessor.ProcessAssets(appid, filePath, batchsize)
}

var batchsize int = 1000 // read from config file

func init() {
	var err error
	batchsize, err = utils.GetBatchSize("../../config/config.json")
	if err != nil {
		log.Println("error in reading config file ", err)
	}
}
