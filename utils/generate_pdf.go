package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"github.com/jung-kurt/gofpdf"
	"github.com/saarthi123/saarthi-backend/models"
)

const logoPath = "assets/logo.png"

func GeneratePDF(tips []models.FinancialTip) []byte {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	addHeader(pdf)
	addTitle(pdf, "Saarthi AI Financial Tips")
	addTipsTable(pdf, tips)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return []byte("Failed to generate PDF")
	}
	return buf.Bytes()
}


func addHeader(pdf *gofpdf.Fpdf) {
	imageType := getImageType(logoPath)
	if imageType != "" {
		pdf.ImageOptions(logoPath, 10, 10, 30, 0, false, gofpdf.ImageOptions{
			ImageType: imageType,
			ReadDpi:   true,
		}, 0, "")
	}

	pdf.SetFont("Arial", "B", 14)
	pdf.SetXY(50, 15)
	pdf.Cell(40, 10, "Aryavarta Saarthi")
	pdf.Ln(20)
}

// Helper to detect image type
func getImageType(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".png":
		return "PNG"
	case ".jpg", ".jpeg":
		return "JPG"
	default:
		return ""
	}
}

func addTitle(pdf *gofpdf.Fpdf, title string) {
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, title, "", 1, "C", false, 0, "")
	pdf.Ln(5)
}

func addTipsTable(pdf *gofpdf.Fpdf, tips []models.FinancialTip) {
	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(200, 220, 255)

	pdf.CellFormat(80, 10, "Title", "1", 0, "C", true, 0, "")
	pdf.CellFormat(110, 10, "Body", "1", 1, "C", true, 0, "")

	for _, tip := range tips {
		pdf.CellFormat(80, 10, tip.Title, "1", 0, "", false, 0, "")
		pdf.MultiCell(110, 10, tip.Body, "1", "", false)
	}
}

// GenerateText creates a plain text version of financial tips
func GenerateText(tips []models.FinancialTip) string {
	var content string
	for _, tip := range tips {
		content += fmt.Sprintf("Title: %s\n%s\n\n", tip.Title, tip.Body)
	}
	return content
}