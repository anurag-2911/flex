package assetprocessor

import (
	"assetmgmt/pkg/model"
	"fmt"
	"strings"
)

// removeDuplicateComputers filters out duplicate computer entries.
func RemoveDuplicateAssets(assets []model.Asset) []model.Asset {
	seen := make(map[string]model.Asset)
	for _, asset := range assets {
		key := fmt.Sprintf("%s-%s", asset.ComputerID, asset.UserID)
		if existingComp, exists := seen[key]; exists {
			// If the existing entry is a desktop and the new one is a laptop, replace it.
			if strings.ToLower(existingComp.ComputerType) != model.LAPTOP && strings.ToLower(asset.ComputerType) == model.LAPTOP {
				seen[key] = asset
			}
		} else {
			seen[key] = asset
		}
	}

	uniqueComputers := make([]model.Asset, 0, len(seen))
	for _, comp := range seen {
		uniqueComputers = append(uniqueComputers, comp)
	}

	return uniqueComputers
}

// calculateMinimumCopies calculates the minimum number of application copies needed.
func CalculateMinimumCopies(assets []model.Asset,assetid string) int {
	userLaptopCount := make(map[string]int)
	userTotalCount := make(map[string]int)

	for _, asset := range assets {
		if asset.ApplicationID == assetid {
			userTotalCount[asset.UserID]++
			if strings.ToLower(asset.ComputerType) == model.LAPTOP {
				userLaptopCount[asset.UserID]++
			}
		}
	}

	totalCopies := 0
	for userID := range userTotalCount {
		if userLaptopCount[userID] > 0 {
			// At least one laptop allows for two installations per copy.
			copies := (userTotalCount[userID] + 1) / 2 // Ceiling division for odd numbers.
			totalCopies += copies
		} else {
			// No laptops mean each desktop requires a separate copy.
			totalCopies += userTotalCount[userID]
		}
	}

	return totalCopies
}