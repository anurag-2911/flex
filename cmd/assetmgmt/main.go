package main

import (
	"assetmgmt/pkg/assetprocessor"
	"assetmgmt/pkg/commandline"
	"assetmgmt/pkg/model"
	"fmt"
)

func main() {
	// Example input data
	assets := []model.Asset{
		{ComputerID: "1", UserID: "1", ApplicationID: "374", ComputerType: "LAPTOP", Comment: ""},
		{ComputerID: "2", UserID: "1", ApplicationID: "374", ComputerType: "DESKTOP", Comment: ""},
		{ComputerID: "3", UserID: "2", ApplicationID: "374", ComputerType: "DESKTOP", Comment: ""},
		{ComputerID: "4", UserID: "2", ApplicationID: "374", ComputerType: "DESKTOP", Comment: ""},
		{ComputerID: "2", UserID: "2", ApplicationID: "374", ComputerType: "desktop", Comment: ""},
	}
	appid := commandline.ReadCommandLine()
	fmt.Println("application id ", appid)
	if appid != "" {
		// Remove duplicates before processing.
		uniqueassets := assetprocessor.RemoveDuplicateAssets(assets)
		copiesNeeded := assetprocessor.CalculateMinimumCopies(uniqueassets, appid)
		fmt.Printf("Minimum copies needed: %d\n", copiesNeeded)
	}
}
