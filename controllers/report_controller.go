package controllers

import (
	"food-ordering/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportService services.ReportService
}

func NewReportController(service services.ReportService) *ReportController {
	return &ReportController{reportService: service}
}
func (c *ReportController) GetSalesReport(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
		return
	}

	orders, err := c.reportService.GetSalesReport()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"sales_report": orders})
}
