package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type SPJHandler struct {
	Service *services.SPJService
}

func NewSPJHandler(service *services.SPJService) *SPJHandler {
	return &SPJHandler{
		Service: service,
	}
}

func (h *SPJHandler) Create(c *gin.Context) {
	var data models.SPJOrder

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "format JSON tidak valid",
			"error":   err.Error(),
		})
		return
	}

	if err := h.Service.Create(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "data SPJ berhasil dibuat",
		"data":    data,
	})
}

func (h *SPJHandler) FindAll(c *gin.Context) {
	data, err := h.Service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal mengambil data SPJ",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *SPJHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID tidak valid",
		})
		return
	}

	data, err := h.Service.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "data SPJ tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *SPJHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID tidak valid",
		})
		return
	}

	var data models.SPJOrder

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "format JSON tidak valid",
			"error":   err.Error(),
		})
		return
	}

	if err := h.Service.Update(uint(id), &data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	updated, err := h.Service.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal mengambil data setelah update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "data SPJ berhasil diupdate",
		"data":    updated,
	})
}

func (h *SPJHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID tidak valid",
		})
		return
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "data SPJ tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "data SPJ berhasil dihapus",
	})
}