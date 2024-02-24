package assetprocessor

import (
	"assetmgmt/pkg/model"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Global map to track the minimum number of copies per user.
var userCopies = make(map[string]int)
var mu sync.Mutex // For safely updating userCopies

func processComputers(computers []model.Asset,appid string) {
	fmt.Println("processing assets")
    localCopies := make(map[string]map[string]bool) // UserID to a map of ComputerType (normalized) to bool
    for _, computer := range computers {
        if computer.ApplicationID != appid {
            continue
        }
        // Normalize ComputerType to lowercase
        computerType := strings.ToLower(computer.ComputerType)
        if _, exists := localCopies[computer.UserID]; !exists {
            localCopies[computer.UserID] = make(map[string]bool)
        }
        // Mark the computer type as present for the user
        localCopies[computer.UserID][computerType] = true
    }

    mu.Lock()
    for userID, types := range localCopies {
        // If the user has at least one laptop, only one copy is required.
        // Otherwise, increment the copies required for each desktop.
        _, hasLaptop := types[model.LAPTOP]
        if hasLaptop {
            userCopies[userID] = max(userCopies[userID], 1)
        } else if _, hasDesktop := types[model.DESKTOP]; hasDesktop {
            // If no laptop but desktops, ensure at least one copy is accounted for, per desktop.
            userCopies[userID] += 1
        }
    }
    mu.Unlock()
}

// Helper function to find the maximum of two integers.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func readAndProcessCSV(appid string,filepath string) error {
	fmt.Println("read and process CSV")
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
			fmt.Println("reached EOF")
            break
        }
        if err != nil {
			fmt.Println("error in reading file ",err)
            return err
        }

        // Process in batches of 1000
        var batch []model.Asset
        for i := 0; i < 1000; i++ {
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
            processComputers(batch,appid)
        }(batch)
    }

    wg.Wait()
    return nil
}

func ProcessAssets(appID string,filePath string){
	err := readAndProcessCSV(appID,filePath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    totalCopies := 0
	fmt.Println("total copies")
    for _, copies := range userCopies {
        totalCopies += copies
    }

    fmt.Printf("Total application copies required: %d\n", totalCopies)
}