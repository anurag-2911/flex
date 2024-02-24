package assetprocessor

import (
	"assetmgmt/pkg/model"
	"math/rand"
	"strconv"
	"testing"
)

// TestNormalizeData tests the normalization of computer types to lowercase.
func TestNormalizeData(t *testing.T) {
	ap := assetProcessor{}
	computers := []model.Asset{
		{ComputerType: "LAPTOP"},
		{ComputerType: "DESKTOP"},
	}
	normalized := ap.NormalizeData(computers)

	for _, comp := range normalized {
		if comp.ComputerType != "laptop" && comp.ComputerType != "desktop" {
			t.Errorf("Expected computer type to be normalized to lowercase, got %s", comp.ComputerType)
		}
	}
}

// TestBusinessLogic tests the business logic for calculating application copies.
func TestBusinessLogic(t *testing.T) {
	ap := assetProcessor{userCopies: make(map[string]int)}
	computers := []model.Asset{
		{UserID: "1", ApplicationID: "374", ComputerType: "laptop"},
		{UserID: "1", ApplicationID: "374", ComputerType: "desktop"},
		// Add more test cases as needed
	}
	ap.BusinessLogic(computers, "374")

	if ap.userCopies["1"] != 1 {
		t.Errorf("Expected 1 copy for user 1, got %d", ap.userCopies["1"])
	}
}

func BenchmarkNormalizeData(b *testing.B) {
	ap := assetProcessor{}
	computers := make([]model.Asset, 1000)
	for i := range computers {
		computers[i] = model.Asset{ComputerType: "LAPTOP"}
	}

	for i := 0; i < b.N; i++ {
		ap.NormalizeData(computers)
	}
}
// Benchmark test for processing a large number of records.
func BenchmarkProcessComputers(b *testing.B) {
	ap := &assetProcessor{userCopies: make(map[string]int)}

	// Example: Generate 1 million records
	assets := generateRandomAssets(1_000_000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ap.processComputers(assets, "374")
	}
}
// Generates a random Asset record.
func generateRandomAsset() model.Asset {
	return model.Asset{
		UserID:        strconv.Itoa(rand.Intn(10000)), // Assuming UserID range
		ApplicationID: "374",                           // Fixed for this example
		ComputerType:  []string{"laptop", "desktop"}[rand.Intn(2)], // Randomly choose
	}
}
// Generates a slice of Asset records.
func generateRandomAssets(n int) []model.Asset {
	var assets []model.Asset
	for i := 0; i < n; i++ {
		assets = append(assets, generateRandomAsset())
	}
	return assets
}

