package controllers

import (
	"encoding/csv"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

// GET /loan/summary
func GetLoanSummary(c *gin.Context) {
	summary := models.LoanSummary{
		ID:            "loan-001",
		AccountNumber: "1234567890",
		Principal:     500000,
		Remaining:     320000,
		InterestRate:  8.5,
		TermMonths:    36,
		StartDate:     "2024-01-15",
	}
	// ✅ Send struct directly
c.JSON(http.StatusOK, gin.H{"summary": summary})
}

// GET /loan/emi-schedule
func GetEMISchedule(c *gin.Context) {
	start, _ := time.Parse("2006-01-02", "2024-01-15")
	emis := []models.EMISchedule{}

	for i := 1; i <= 36; i++ {
		due := start.AddDate(0, i-1, 0)
		emis = append(emis, models.EMISchedule{
			Month:        due.Format("Jan 2006"),
			DueDate:      due.Format("2006-01-02"),
			Principal:    13000,
			Interest:     1800,
			TotalPayment: 14800,
			Status: func() string {
				if due.Before(time.Now()) {
					return "Paid"
				}
				return "Due"
			}(),
		})
	}

	// ✅ Send slice directly
c.JSON(http.StatusOK, gin.H{"schedule": emis})
}

// GET /loan/schedule/download
func DownloadScheduleCSV(c *gin.Context) {
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=emi-schedule.csv")
	c.Writer.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	writer.Write([]string{"Month", "Due Date", "Principal", "Interest", "Total", "Status"})

	start, _ := time.Parse("2006-01-02", "2024-01-15")
	for i := 1; i <= 36; i++ {
		due := start.AddDate(0, i-1, 0)
		status := "Due"
		if due.Before(time.Now()) {
			status = "Paid"
		}
		writer.Write([]string{
			due.Format("Jan 2006"),
			due.Format("2006-01-02"),
			"13000",
			"1800",
			"14800",
			status,
		})
	}
}
