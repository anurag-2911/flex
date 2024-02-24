package main

import (
	"assetmgmt/pkg/assetprocessor"
	"assetmgmt/pkg/commandline"
	"log"
)


func main() {
   appid,filePath:=commandline.GetCommandLineArguments()
   log.Println("application ID and filepath : ",appid,filePath)
   assetprocessor.ProcessAssets(appid,filePath)
}
