package tests

import (
	"os"
	"path/filepath"
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
	PathFileformat TestPathFileformat `yaml:"path-fileformat"`
	Time           string             `yaml:"time"`
	Auto           bool               `yaml:"auto"`
}

type TestPathFileformat struct {
	Include []string `yaml:"include"`
	Exclude []string `yaml:"exclude"`
}

func TestRemoveFile(t *testing.T) {
	//file, _ := os.ReadFile("./test.yaml")
	//
	//var config TestConfig
	//_ = yaml.Unmarshal(file, &config)
}

func removeFile(path string, excludePath []string) error {
	return filepath.Walk(path, func(file string, info os.FileInfo, err error) error {

	})
}
