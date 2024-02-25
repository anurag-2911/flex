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

func TestGetBatchSizeNonExistentFile(t *testing.T) {
	// Execute: Call GetBatchSize with a non-existent file path.
	batchSize, err := GetBatchSize("non-existent-file.json")

	// Assert: Expect error and default batch size.
	if err == nil {
		t.Errorf("Expected error for non-existent file")
	}
	if batchSize != defaultBatchSize {
		t.Errorf("Expected default batch size (%d), got %d", defaultBatchSize, batchSize)
	}
}

func TestGetBatchSizeInvalidJSON(t *testing.T) {
	// Setup: Create a temporary config file with invalid JSON.
	config := `{"invalid_json_format": 100}`
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

	// Assert: Expect error and default batch size.
	if err != nil {
		t.Errorf("Expected error for invalid JSON")
	}
	if batchSize == defaultBatchSize {
		t.Errorf("Expected default batch size (%d), got %d", defaultBatchSize, batchSize)
	}
}



