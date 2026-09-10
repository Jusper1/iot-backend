package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type InaprocHandler struct{
	Service *services.InaprocService
}

func NewInaprocHandler (
	service *services.InaprocService,
)*InaprocHandler {
	return &InaprocHandler{
		Service: service,
	}
}

func (h *InaprocHandler) Create(c *gin.Context) {
	var data models.InaprocOrder

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message":"Format JSON tidak valid",
			"error":err.Error(),
		})
		return
	} 

	err := h.Service.Create(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":true,
		"message":"Data Berhasil Ditambahkan",
		"data":data,
	})
}

func (h *InaprocHandler) FindAll(c *gin.Context){

	data, err := h.Service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":false,
			"message":"Gagal Mengambil Data Inaproc",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"message":"Data Inaproc berhasil Diambil",
		"data":	data,
	})
}

func (h *InaprocHandler) FindByID(c *gin.Context){

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"ID Tidak Valid",
		})
		return
	}

	data, err := h.Service.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success":false,
			"message":"Data Inaproc tidak berhasil ditemukan",
			"data":data,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"data":data,
	})
}

func (h *InaprocHandler) Update(c *gin.Context){

	idParam := c.Param("id")

	id,err :=strconv.ParseUint(idParam, 10, 64)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"ID tidak valid",	
		})
	}

	var data models.InaprocOrder

	if err := c.ShouldBindJSON(&data); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"Format JSON Tidak Valid",
			"error":err.Error(),
		})
		return
	}

	err = h.Service.Update(uint(id), &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":err.Error(),
		})
		return
	}

	updatedData, err := h.Service.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":false,
			"message":"gagal mengambil data setelah Update",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"message":"Data inaproc berhasil di update",
		"data":updatedData,
	})
}

func (h *InaprocHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"ID tidak valid",
		})
		return
	}

	err = h.Service.Delete(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success":false,
			"message":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"message":"Data berhasil di hapus",
	})
}