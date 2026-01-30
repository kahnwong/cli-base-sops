package cli_base_sops

import (
	"os"
	"path/filepath"
	"testing"
)

// TestConfig is a simple test struct for unmarshalling
type TestConfig struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func TestReadYamlSops_FileNotFound(t *testing.T) {
	// Test with non-existent file
	_, err := ReadYamlSops[TestConfig]("/nonexistent/path/config.yaml")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestReadYamlSops_EmptyPath(t *testing.T) {
	// Test with empty path
	_, err := ReadYamlSops[TestConfig]("")
	if err == nil {
		t.Error("Expected error for empty path, got nil")
	}
}

func TestDecryptSops_FileNotFound(t *testing.T) {
	// Test decryptSops with non-existent file
	_, err := decryptSops("/nonexistent/file.yaml", "yaml")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestDecryptSops_InvalidFormat(t *testing.T) {
	// Create a temporary non-SOPS file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.yaml")

	content := []byte("name: test\nvalue: data")
	err := os.WriteFile(tmpFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Try to decrypt a non-SOPS file (should fail)
	_, err = decryptSops(tmpFile, "yaml")
	if err == nil {
		t.Error("Expected error for non-SOPS file, got nil")
	}
}
