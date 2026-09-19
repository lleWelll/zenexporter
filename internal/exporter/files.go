package exporter

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	exportedCsvNameTemplate = "TransactionHistory-*.csv"
)

var outDirPath string

func createDirIfNotExists(args map[string]any) error {
	raw, ok := args[outDirFlag]
	if !ok || raw == nil {
		return ErrInvalidOutputDir
	}

	path, ok := raw.(string)
	if !ok {
		return ErrInvalidOutputDir
	}

	normalised := normalisePath(path)
	if err := os.MkdirAll(normalised, 0o755); err != nil {
		return fmt.Errorf("create output dir %s: %w", normalised, err)
	}

	outDirPath = normalised
	return nil
}

func normalisePath(path string) string {
	normalised := strings.TrimSpace(path)
	normalised = strings.TrimSuffix(normalised, "/")

	return normalised
}

func createNewCsv(data []byte) (string, error) {
	fileName := fmt.Sprintf(exportedCsvNameTemplate, time.Now())
	path := outDirPath + "/" + fileName

	csv, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer csv.Close()

	if _, err := csv.Write(data); err != nil {
		return "", err
	}

	return path, nil
}
