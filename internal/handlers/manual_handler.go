package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type ManualHandler struct{
	Service *services.ManualService
}

func NewManualHandler (
	service *services.ManualService,
)*ManualHandler{
	return &ManualHandler{
		Service: service,
	}
}

func (h *ManualHandler) Create(c *gin.Context) {
	var order models.ManualOrder

	if err := c.ShouldBindJSON(&order); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
        "success": false,
        "message": "format JSON tidak valid",
        "error": err.Error(),
    })
    return
	}
	
	if err := h.Service.Create(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "pemesanan manual berhasil dibuat",
		"data":    order,
	})
}

func (h *ManualHandler) GetAll(c *gin.Context) {
	orders, err := h.Service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal mengambil data pemesanan manual",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    orders,
	})
}

func (h *ManualHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID tidak valid",
		})
		return
	}

	order, err := h.Service.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "pemesanan manual tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal mengambil data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    order,
	})
}

func (h *ManualHandler) Update(c *gin.Context) {
	var order models.ManualOrder

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID tidak valid",
		})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message":"format JSON tidak Valid",
			"error": err.Error(),
		})
	}

	if err := h.Service.Update(id, &order); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "pemesanan manual tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "pemesanan manual berhasil diupdate",
		"data":    order,
	})
}

func (h *ManualHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID tidak valid",
		})
		return
	}

	if err := h.Service.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "pemesanan manual tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal menghapus data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "pemesanan manual berhasil dihapus",
	})
}