package exporter

import (
	"encoding/json"

	"github.com/Sarthak-Kamble/ripper/pkg/models"
)

type JSONExporter struct{}

func NewJSONExporter() *JSONExporter {
	return &JSONExporter{}
}

func (e *JSONExporter) Export(
	page models.Page,
) ([]byte, error) {

	return json.MarshalIndent(
		page,
		"",
		"  ",
	)
}
