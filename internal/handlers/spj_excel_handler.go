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

type SPJExcelHandler struct {
	Service *services.SPJService
}

func NewSPJExcelHandler(
	service *services.SPJService,
) *SPJExcelHandler {
	return &SPJExcelHandler{
		Service: service,
	}
}

func (h *SPJExcelHandler) Export(c *gin.Context) {

	orders, err := h.Service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	file := excelize.NewFile()

	sheet := "SPJ"

	file.SetSheetName("Sheet1", sheet)

	headers := []string{
		"NO",
		"TGL Print",
		"Tanggal Paraf",
		"Tanggal Sign",
		"Tanggal Kirim",
		"Nama Dinkes/PKM",
		"Kertas",
		"Kebutuhan SPJ",
		"Jumlah Rangkap",
		"Jenis File",
		"UP, Nomor Telepon dan Alamat",
		"Status Pengiriman",
		"Keterangan",
		"PIC Print",
	}

	for i, header := range headers {

		cell, _ := excelize.CoordinatesToCellName(
			i+1,
			1,
		)

		file.SetCellValue(
			sheet,
			cell,
			header,
		)
	}

	for rowIndex, order := range orders {

		row := rowIndex + 2

		file.SetCellValue(
			sheet,
			fmt.Sprintf("A%d", row),
			rowIndex+1,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("B%d", row),
			order.TglPrint,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("C%d", row),
			order.TglParaf,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("D%d", row),
			order.TglSign,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("E%d", row),
			order.TglKirim,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("F%d", row),
			order.NamaDinkesPKM,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("G%d", row),
			order.Kertas,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("H%d", row),
			order.KebutuhanSPJ,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("I%d", row),
			order.JumlahRangkap,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("J%d", row),
			order.JenisFile,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("K%d", row),
			order.UPNomorTeleponAlamat,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("L%d", row),
			order.StatusPengiriman,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("M%d", row),
			order.Keterangan,
		)

		file.SetCellValue(
			sheet,
			fmt.Sprintf("N%d", row),
			order.PICPrint,
		)
	}

	for i := range headers {

		cell, _ := excelize.CoordinatesToCellName(
			i+1,
			1,
		)

		file.SetColWidth(
			sheet,
			cell[:1],
			cell[:1],
			20,
		)
	}

	c.Header(
		"Content-Disposition",
		`attachment; filename="spj.xlsx"`,
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

func (h *SPJExcelHandler) Import(c *gin.Context) {

	fileHeader, err := c.FormFile("file")

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "file Excel wajib diupload",
		})

		return
	}

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

		order := models.SPJOrder{}

		order.TglPrint = models.Date(
			getExcelString(row, 1),
		)

		order.TglParaf = models.Date(
			getExcelString(row, 2),
		)

		order.TglSign = models.Date(
			getExcelString(row, 3),
		)

		order.TglKirim = models.Date(
			getExcelString(row, 4),
		)

		order.NamaDinkesPKM = getExcelString(
			row,
			5,
		)

		order.Kertas = getExcelString(
			row,
			6,
		)

		order.KebutuhanSPJ = getExcelString(
			row,
			7,
		)

		order.JumlahRangkap = uint(
			getExcelInt(row, 8),
		)

		order.JenisFile = getExcelString(
			row,
			9,
		)

		order.UPNomorTeleponAlamat = getExcelString(
			row,
			10,
		)

		order.StatusPengiriman = getExcelString(
			row,
			11,
		)

		order.Keterangan = getExcelString(
			row,
			12,
		)

		order.PICPrint = getExcelString(
			row,
			13,
		)

		if order.NamaDinkesPKM == "" {

			errorsImport = append(
				errorsImport,
				fmt.Sprintf(
					"baris %d: Nama Dinkes/PKM wajib diisi",
					excelRow,
				),
			)

			continue
		}

		if err := h.Service.Create(
			&order,
		); err != nil {

			errorsImport = append(
				errorsImport,
				fmt.Sprintf(
					"baris %d (%s): %s",
					excelRow,
					order.NamaDinkesPKM,
					err.Error(),
				),
			)

			continue
		}

		imported++
	}

	c.JSON(http.StatusOK, gin.H{

		"success": true,

		"message": "import Excel selesai",

		"imported": imported,

		"errors": errorsImport,
	})
}

func getSPJExcelString(
	row []string,
	index int,
) string {

	if index >= len(row) {
		return ""
	}

	return strings.TrimSpace(
		row[index],
	)
}

func getSPJExcelInt(
	row []string,
	index int,
) int64 {

	value := getSPJExcelString(
		row,
		index,
	)

	if value == "" {
		return 0
	}

	number, err := strconv.ParseInt(
		value,
		10,
		64,
	)

	if err != nil {

		floatValue, err := strconv.ParseFloat(
			value,
			64,
		)

		if err != nil {
			return 0
		}

		return int64(floatValue)
	}

	return number
}