package service

import (
	"bytes"
	"context"
	"farms-planning/app/repo"
	"farms-planning/pkg/runapp/fiber/logger"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type other struct {
	log logger.LoggerFiber
	r   repo.Repo
}

func NewServices(log logger.LoggerFiber, r repo.Repo) Services {
	return &other{
		log: log,
		r:   r,
	}
}

// ################### OTHER ###################
func (s *other) ProvinceGetAll(ctx context.Context) (result interface{}, err error) {
	return s.r.ProvinceGetAll()
}

func (s *other) GeneratePDF(ctx context.Context, userData map[string]interface{}) ([]byte, error) {

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a page
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "B", 16)

	// Add content to the PDF
	for key, value := range userData {
		line := fmt.Sprintf("%s: %v", key, value)
		pdf.Cell(0, 10, line)
		pdf.Ln(10)
	}

	// Generate PDF
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
