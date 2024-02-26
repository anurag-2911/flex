package assetprocessor

import (
	"assetmgmt/pkg/model"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

/*
A struct that encapsulates the state and functionality for processing asset data,
including counting user copies of software applications.
*/
type assetProcessor struct {
	userCopies map[string]int
	mu         sync.Mutex
}

/*
This method takes a slice of model.Asset representing computer assets,
normalizes the ComputerType field to lowercase for consistency,
and returns the normalized slice.
*/
func (ap *assetProcessor) NormalizeData(computers []model.Asset) []model.Asset {
	var normalized []model.Asset
	for _, comp := range computers {
		comp.ComputerType = strings.ToLower(comp.ComputerType)
		normalized = append(normalized, comp)
	}
	return normalized
}

/*
Processes a slice of normalized model.Asset, filtering by the given appid.
It calculates the required number of application copies based on the business rules:
each user needs only one copy if they have a laptop, otherwise one copy per desktop.
It updates the userCopies map with the total counts.
*/
func (ap *assetProcessor) BusinessLogic(computers []model.Asset, appid string) {
	localCopies := make(map[string]map[string]int)
	for _, computer := range computers {
		if computer.ApplicationID != appid {
			continue
		}
		if _, exists := localCopies[computer.UserID]; !exists {
			localCopies[computer.UserID] = make(map[string]int)
		}
		// Increment the count for either computer type
		localCopies[computer.UserID][computer.ComputerType] += 1
	}

	ap.mu.Lock()
	defer ap.mu.Unlock()
	for userID, types := range localCopies {
		_, hasLaptop := types[model.LAPTOP]
		_, hasDesktop := types[model.DESKTOP]
		if hasLaptop {
			// If the user has at least one laptop, they need only one copy
			ap.userCopies[userID] = max(ap.userCopies[userID], 1)
		} else if hasDesktop {
			// If the user has only desktops, they can have up to two copies
			// This assumes the business rule allows for 2 copies if only desktops are present
			desktopCount := types[model.DESKTOP]
			if desktopCount > 1 { // Check if there are more than one desktop
				ap.userCopies[userID] = max(ap.userCopies[userID], 2)
			} else {
				ap.userCopies[userID] = max(ap.userCopies[userID], desktopCount)
			}
		}
	}
}

// Orchestrates the processing of computer assets for a specific application ID.
// It first normalizes the data and then applies the business logic to determine the necessary
// application copies
func (ap *assetProcessor) processComputers(computers []model.Asset, appid string) {
	normalizedComputers := ap.NormalizeData(computers)
	ap.BusinessLogic(normalizedComputers, appid)
}

/*
Reads asset data from a CSV file at filepath, processing the data in batches of size batchsize.
Each batch is processed concurrently, utilizing goroutines to handle the normalization and business logic.
Errors during file reading or processing are returned to the caller.
*/
func (ap *assetProcessor) readAndProcessCSV(appid string, filepath string, batchsize int) error {
	log.Println("read and process CSV")
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	var wg sync.WaitGroup

	for {
		records, err := csvReader.Read()
		if err == io.EOF {
			log.Println("reached EOF")
			break
		}
		if err != nil {
			log.Println("error in reading file ", err)
			return err
		}

		// Process in batches as defined in config file
		var batch []model.Asset
		for i := 0; i < batchsize; i++ {
			if len(records) == 0 { // If no more records, break
				break
			}
			if len(records) < 4 { // Skip if record is incomplete
				continue
			}
			batch = append(batch, model.Asset{
				UserID:        records[1],
				ApplicationID: records[2],
				ComputerType:  records[3],
			})
			records, err = csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
		}

		wg.Add(1)
		go func(batch []model.Asset) {
			defer wg.Done()
			ap.processComputers(batch, appid)
		}(batch)
	}

	wg.Wait()
	return nil
}

// The entry point for processing assets from a CSV file. It initializes an assetProcessor,
// reads and processes the CSV file, and finally calculates and displays the total number of application
// copies required.
func ProcessAssets(appID string, filePath string, batchsize int) {
	ap := &assetProcessor{userCopies: make(map[string]int)}
	err := ap.readAndProcessCSV(appID, filePath, batchsize)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	totalCopies := 0
	log.Println("total copies")
	for _, copies := range ap.userCopies {
		totalCopies += copies
	}

	log.Printf("Total application copies required: %d\n", totalCopies)
}

// A helper function that returns the maximum of two integers.
// Used to ensure the user copy count does not decrease if a user already has more copies
// than a new calculation suggests.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
