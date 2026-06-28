package exporter

import (
	"testing"

	"github.com/Sarthak-Kamble/ripper/pkg/models"
)

func TestJSONExporterExport(t *testing.T) {

	exporter := NewJSONExporter()

	page := models.Page{
		URL:    "https://example.com",
		Title:  "Example",
		Status: "200 OK",
	}

	data, err := exporter.Export(page)

	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 0 {
		t.Fatal("expected json output")
	}
}
