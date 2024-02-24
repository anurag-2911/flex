package utils

import (	
	"os"
	"testing"
)

// TestGetBatchSizeValidConfig tests GetBatchSize with a valid configuration file.
func TestGetBatchSizeValidConfig(t *testing.T) {
	// Setup: Create a temporary config file with valid JSON.
	config := `{"batchsize": 500}`
	tmpfile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up after the test.

	if _, err := tmpfile.Write([]byte(config)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Execute: Call GetBatchSize with the path to the temp file.
	batchSize, err := GetBatchSize(tmpfile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if batchSize != 500 {
		t.Errorf("Expected batch size 500, got %d", batchSize)
	}
}



