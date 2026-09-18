package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iot-backend/internal/services"
)

type DashboardHandler struct {
	Service *services.DashboardService
}

func NewDashboardHandler(
	service *services.DashboardService,
) *DashboardHandler {
	return &DashboardHandler{
		Service: service,
	}
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {

	periode := c.DefaultQuery(
		"periode",
		"month",
	)

	if periode != "day" &&
		periode != "month" &&
		periode != "year" {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "periode harus day, month, atau year",
		})

		return
	}

	data, err := h.Service.GetDashboard(
		periode,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal mengambil data dashboard",
			"error":   err.Error(),
		})

		return
	}

	var totalInaproc int64
	var totalManual int64
	var totalAlat int64

	var totalInaprocDibayar int64
	var totalManualDibayar int64
	var totalAlatDibayar int64

	for _, item := range data {

		totalInaproc += item.Inaproc
		totalManual += item.Manual
		totalAlat += item.Alat

		totalInaprocDibayar += item.InaprocDibayar
		totalManualDibayar += item.ManualDibayar
		totalAlatDibayar += item.AlatDibayar
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"periode": periode,

		"data": gin.H{

			"total_pembelian": gin.H{
				"inaproc": totalInaproc,
				"manual":  totalManual,
				"alat":    totalAlat,
			},

			"total_dibayar": gin.H{
				"inaproc": totalInaprocDibayar,
				"manual":  totalManualDibayar,
				"alat":    totalAlatDibayar,
			},

			"grafik": data,
		},
	})
}