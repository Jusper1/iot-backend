package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type ManualExcelHandler struct{
	Service *services.ManualService
}

func NewManualExcelHandler(
	service *services.ManualService,
)*ManualExcelHandler{
	return &ManualExcelHandler{
		Service: service,
	}
}

func (h *ManualExcelHandler) Export(c *gin.Context) {
	orders, err := h.Service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":false,
			"message":err.Error(),
		})
		return
	}

	file :=excelize.NewFile()
	sheet := "Pemesanan Manual"

	file.SetSheetName("Sheet1",sheet)

	headers := []string{
		"KODE",
		"DINKES/PUSKESMAS",
		"EMAIL",
		"STATUS ODOO",
		"STATUS",
		"QTY",
		"WAKTU PENGIRIMAN (BULAN)",
		"NAMA PIC",
		"NIK",
		"NPWP",
		"NO PO",
		"TANGGAL PO",
		"NO BAST",
		"TANGGAL BAST",
		"HARGA + PPN",
		"TANGGAL UANG MASUK",
		"JUMLAH UANG MASUK",
		"PERIODE LANGGANAN",
		"NOMOR INVOICE",
		"ALAMAT",
		"KOTA/KAB",
	}

	for i,header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		file.SetCellValue(sheet, cell, header)
	}

	for rowIndex, order := range orders {
		row := rowIndex + 2

		file.SetCellValue(sheet, fmt.Sprintf("A%d", row), order.Kode)
		file.SetCellValue(sheet, fmt.Sprintf("B%d", row), order.DinkesPuskesmas)
		file.SetCellValue(sheet, fmt.Sprintf("C%d", row), order.Email)
		file.SetCellValue(sheet, fmt.Sprintf("D%d", row), order.StatusOdoo)
		file.SetCellValue(sheet, fmt.Sprintf("E%d", row), order.Status)
		file.SetCellValue(sheet, fmt.Sprintf("F%d", row), order.Qty)

		if order.WaktuPengirimanBulan != nil{
			file.SetCellValue(
				sheet,
				fmt.Sprintf("G%d", row),
				order.WaktuPengirimanBulan,
			)
		}

		file.SetCellValue(sheet, fmt.Sprintf("H%d", row),order.NamaPIC)
		file.SetCellValue(sheet, fmt.Sprintf("I%d", row),order.NIK)
		file.SetCellValue(sheet, fmt.Sprintf("J%d", row),order.NPWP)
		file.SetCellValue(sheet, fmt.Sprintf("K%d", row),order.NoPO)
		file.SetCellValue(sheet, fmt.Sprintf("L%d", row),order.TglPO)
		file.SetCellValue(sheet, fmt.Sprintf("M%d", row),order.NoBAST)
		file.SetCellValue(sheet, fmt.Sprintf("N%d", row),order.TglBAST)
		file.SetCellValue(sheet, fmt.Sprintf("O%d", row),order.HargaPPN)
		file.SetCellValue(sheet, fmt.Sprintf("P%d", row),string(order.TglUangMasuk))
		file.SetCellValue(sheet,fmt.Sprintf("Q%d", row),order.JumlahUangMasuk)
		file.SetCellValue(sheet,fmt.Sprintf("R%d", row),order.PeriodeLangganan)
		file.SetCellValue(sheet,fmt.Sprintf("S%d", row),order.NomorInvoice)
		file.SetCellValue(sheet,fmt.Sprintf("T%d", row),order.Alamat)
		file.SetCellValue(sheet,fmt.Sprintf("U%d", row),order.KotaKab)
	}

	for i := range headers{
		cell, _ :=excelize.CoordinatesToCellName(i+1, 1)
		file.SetColWidth(sheet, cell[:1], cell[:1], 20)
	}

	c.Header(
		"Content-Disposition",
		`attachment; filename="pemesanan_manual.xlsx"`,
	)
	c.Header(
		"Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	)

	if err := file.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
}

func (h *ManualExcelHandler) Import(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "file Excel wajib diupload",
			// "error":   err.Error(),
		})
		return
	}


	// fmt.Println("FILE:", fileHeader.Filename)
	// fmt.Println("SIZE:", fileHeader.Size)

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "file tidak dapat dibuka",
			"error":   err.Error(),
		})
		return
	}
	defer file.Close()

	excel, err := excelize.OpenReader(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "file bukan Excel yang valid",
			"error":   err.Error(),
		})
		return
	}
	defer excel.Close()

	sheet := excel.GetSheetName(0)

	rows, err := excel.GetRows(sheet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal membaca isi Excel",
			"error":   err.Error(),
		})
		return
	}

	if len(rows) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Excel tidak memiliki data",
		})
		return
	}

	var imported int
	var errorsImport []string

	for i, row := range rows[1:] {
		excelRow := i + 2

		if len(row) == 0 {
			continue
		}

		order := models.ManualOrder{}

		order.Kode = getExcelString(row, 0)
		order.DinkesPuskesmas = getExcelString(row, 1)
		order.Email = getExcelString(row, 2)
		order.StatusOdoo = getExcelString(row, 3)
		order.Status = getExcelString(row, 4)

		order.Qty = uint(getExcelInt(row, 5))

		if value := getExcelInt(row, 6); value > 0 {
			v := uint(value)
			order.WaktuPengirimanBulan = &v
		}

		order.NamaPIC = getExcelString(row, 7)
		order.NIK = getExcelString(row, 8)
		order.NPWP = getExcelString(row, 9)
		order.NoPO = getExcelString(row, 10)

		order.TglPO = models.Date(getExcelString(row, 11))

		order.NoBAST = getExcelString(row, 12)
		order.TglBAST = models.Date(getExcelString(row, 13))

		order.HargaPPN = getExcelFloat(row, 14)

		order.TglUangMasuk = models.Date(getExcelString(row, 15))

		order.JumlahUangMasuk = getExcelFloat(row, 16)

		order.PeriodeLangganan = getExcelString(row, 17)
		order.NomorInvoice = getExcelString(row, 18)
		order.Alamat = getExcelString(row, 19)
		order.KotaKab = getExcelString(row, 20)

		if order.Kode == "" {
			errorsImport = append(
				errorsImport,
				fmt.Sprintf("baris %d: KODE wajib diisi", excelRow),
			)
			continue
		}

		if err := h.Service.Create(&order); err != nil {
			errorsImport = append(
				errorsImport,
				fmt.Sprintf(
					"baris %d (%s): %s",
					excelRow,
					order.Kode,
					err.Error(),
				),
			)
			continue
		}

		imported++
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "import Excel selesai",
		"imported": imported,
		"errors":   errorsImport,
	})
}

func getExcelString(row []string, index int) string {
	if index >= len(row) {
		return ""
	}

	return strings.TrimSpace(row[index])
}

func getExcelInt(row []string, index int) int64 {
	value := getExcelString(row, index)

	if value == "" {
		return 0
	}

	number, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0
		}

		return int64(floatValue)
	}

	return number
}

func getExcelFloat(row []string, index int) float64 {
	value := getExcelString(row, index)

	if value == "" {
		return 0
	}

	value = strings.ReplaceAll(value, ",", "")

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return number
}