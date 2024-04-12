package tests

import (
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

type TestConfig struct {
	Monitoring  TestMonitoring `yaml:"monitoring,omitempty"`
	CleanPolicy TestPolicy     `yaml:"clean-policy"`
}

type TestMonitoring struct {
	AlertUsage int    `yaml:"alert-usage"`
	Platform   string `yaml:"platform"`
	ApiKey     string `yaml:"api-key"`
}

type TestPolicy struct {
	Usage          int                `yaml:"usage"`
	PathFileFormat TestPathFileFormat `yaml:"path-file-format"`
	Time           string             `yaml:"time"`
	Auto           bool               `yaml:"auto"`
}

type TestPathFileFormat struct {
	Include []string `yaml:"include"`
	Exclude []string `yaml:"exclude"`
}

func TestYaml(t *testing.T) {
	// Given
	file, _ := os.ReadFile("./test.yaml")

	var config TestConfig
	err := yaml.Unmarshal(file, &config)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// When Then
	if config.Monitoring.AlertUsage != 80 {
		t.Errorf("Expected 80, got %d", config.Monitoring.AlertUsage)
	}
}
