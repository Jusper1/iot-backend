package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type InaprocExcelHandler struct {
	Service * services.InaprocService
}

func NewInaprocExcelHandler(
	service *services.InaprocService,
) *InaprocExcelHandler {
	return &InaprocExcelHandler{
		Service: service,
	}
}

func (h *InaprocExcelHandler) Export(c *gin.Context) {
	
	data, err := h.Service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"message":"Gagal mengambil data inaproc",
			"error": err.Error(),
		})
		return
	}

	file := excelize.NewFile()
	defer file.Close()

	sheetName := "Inaproc"

	index, err := file.NewSheet(sheetName)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message":"gagal mengambil sheet excel",
			"error": err.Error(),
		})
		return
	}

	file.DeleteSheet("Sheet1")
	file.SetActiveSheet(index)

	headers := []string{
		"KODE",
		"NAMA PIC",
		"NIK",
		"PUSKESMAS",
		"STATUS ODOO",
		"NO NPWP",
		"QTY",
		"STATUS",
		"HARGA + PPN",
		"TGL PESANAN",
		"TGL BAST",
		"NO BAST",
		"NO INVOICE KUT",
		"NSFP",
		"JUMLAH UANG MASUK",
		"TGL UANG MASUK",
		"REKENING",
		"NO INVOICE INAPROC",
		"KODE BAYAR",
		"KETERANGAN",
		"ALAMAT",
		"PROVINSI",
	}

	for col, header := range headers {
		cell, err := excelize.CoordinatesToCellName(col+1, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"success":false,
				"message":"gagal membuat cell excel",
				"error":err.Error(),
			})
			return
		}
		
		if err := file.SetCellValue(sheetName, cell ,header); err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"success":false,
				"message":"gagal menulis header excel",
				"error":err.Error(),
			})
			return
		}
	}

	for rowIndex, item := range data {

		row := rowIndex + 2

		values := []interface{}{
			item.Kode,
			item.NamaPic,
			item.NIK,
			item.Puskesmas,
			item.StatusOdoo,
			item.NoNPWP,
			item.Qty,
			item.Status,
			item.HargaPPN,
			string(item.TglPesanan),
			string(item.TglBAST),
			item.NoBAST,
			item.NoInvoiceKUT,
			item.NSFP,
			item.JumlahUangMasuk,
			string(item.TglUangMasuk),
			item.Rekening,
			item.NoInvoiceInaproc,
			item.KodeBayar,
			item.Keterangan,
			item.Alamat,
			item.Provinsi,
		}

		for col, value := range values {
			
			cell, err := excelize.CoordinatesToCellName(col+1, row)

			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{
					"success":false,
					"message":"Gagal membuat cel Excel",
					"error": err.Error(),
				})
				return
			}

			if err := file.SetCellValue(sheetName, cell , value); err != nil{
				c.JSON(http.StatusInternalServerError,gin.H{
					"success":false,
					"message":"Gagal menulis data excel",
					"error":err.Error(),
				})
				return
			}
		}
	}

	headerStyle,err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical: "center",
		},
	})

	if err == nil {
		_ = file.SetCellStyle(
			sheetName,
			"A1",
			"V1",
			headerStyle,
		)
	} 

	_ = file.SetPanes(sheetName, &excelize.Panes{
		Freeze: true,
		Split: false,
		XSplit: 0,
		YSplit: 1,
		TopLeftCell: "A2",
		ActivePane: "bottomLeft",
	})

	columnWidths := map[string]float64{
		"A":  15,
		"B":  20,
		"C":  18,
		"D":  25,
		"E":  18,
		"F":  18,
		"G":  10,
		"H":  15,
		"I":  18,
		"J":  15,
		"K":  15,
		"L":  18,
		"M":  20,
		"N":  18,
		"O":  22,
		"P":  18,
		"Q":  20,
		"R":  25,
		"S":  18,
		"T":  30,
		"U":  30,
		"V":  20,
	}

	for column, width := range columnWidths{
		_ = file.SetColWidth(sheetName, column,column,width)
	}

	filename := fmt.Sprintf(
		"inaproc_%s.xlsx",
		time.Now().Format("20060102_150405"),
	)

	c.Header(
		"Content-Disposition",
		fmt.Sprintf(`attachment; filename="%s"`, filename),
	)

	c.Header(
		"Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	)

	if err := file.Write(c.Writer); err != nil{
		return
	}
}

func (h *InaprocExcelHandler) Import(c *gin.Context) {

	fileHeader, err := c.FormFile("file")

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"file Excel wajib di upload",
			"error":err.Error(),
		})
		return
	}

	filename := strings.ToLower(fileHeader.Filename)

	if !strings.HasSuffix(filename, ".xlsx") {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"file harus berformat .xlsx",
		})
		return
	}

	const maxFileSize = 10 << 20

	if fileHeader.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"file maxsimal 10 mb",
		})
		return
	}

	uploadedFile, err := fileHeader.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":false,
			"message":"gagal membuka file Excel",
			"error":err.Error(),
		})
		return
	}

	defer uploadedFile.Close()

	excel, err := excelize.OpenReader(uploadedFile)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"file excel tidak dapat dibaca",
			"error":err.Error(),
		})
		return
	}

	defer excel.Close()

	sheetList := excel.GetSheetList()

	if len(sheetList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"Excel tidak muncul sheet",
		})
		return
	}

	sheetName := sheetList[0]

	rows, err := excel.GetRows(sheetName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"Gagal membaca baris excel",
			"error":err.Error(),
		})
		return
	}

	if len(rows) <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"excel tidak memiliki data",
		})
		return
	}

	expectedHeaders := []string{
		"KODE",
		"NAMA PIC",
		"NIK",
		"PUSKESMAS",
		"STATUS ODOO",
		"NO NPWP",
		"QTY",
		"STATUS",
		"HARGA + PPN",
		"TGL PESANAN",
		"TGL BAST",
		"NO BAST",
		"NO INVOICE KUT",
		"NSFP",
		"JUMLAH UANG MASUK",
		"TGL UANG MASUK",
		"REKENING",
		"NO INVOICE INAPROC",
		"KODE BAYAR",
		"KETERANGAN",
		"ALAMAT",
		"PROVINSI",
	}

	if len(rows[0]) < len(expectedHeaders) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"jumlah kolom excel tidak seusai",
		})
		return
	}

	for i, expected := range expectedHeaders{
		
		actual := strings.TrimSpace(rows[0][i])

		if strings.ToUpper(actual) != expected {
			c.JSON(http.StatusBadRequest, gin.H{
				"success":false,
				"message":fmt.Sprintf(
					"Header kolom ke-%d tidak sesuai. harus '%s', ditemukan  '%s'",
					i+1,
					expected,
					actual,
				),
			})
			return
		}
	}
	
	var imported []models.InaprocOrder

	var errorsImport []string

	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {

		row := rows [rowIndex]

		if isEmptyExcelRow(row){
			continue
		}

		if len(row) < len(expectedHeaders){
			
			errorsImport = append(
				errorsImport, 
				fmt.Sprintf(
					"Baris %d: jumlah kolom tidak lengkap",
					rowIndex+1,
				),
			)
			continue
		}

		item := models.InaprocOrder{
			Kode: strings.TrimSpace(row[0]),
			NamaPic: strings.TrimSpace(row[1]),
			NIK: strings.TrimSpace(row[2]),
			Puskesmas: strings.TrimSpace(row[3]),
			StatusOdoo: strings.TrimSpace(row[4]),
			NoNPWP: strings.TrimSpace(row[5]),
			Status: strings.TrimSpace(row[7]),
			NoBAST: strings.TrimSpace(row[11]),
			NoInvoiceKUT: strings.TrimSpace(row[12]),
			NSFP: strings.TrimSpace(row[13]),
			Rekening: strings.TrimSpace(row[16]),
			NoInvoiceInaproc: strings.TrimSpace(row[17]),
			KodeBayar: strings.TrimSpace(row[18]),
			Keterangan: strings.TrimSpace(row[19]),
			Alamat: strings.TrimSpace(row[20]),
			Provinsi: strings.TrimSpace(row[21]),
		}

		if item.Kode == "" {
			errorsImport = append(
				errorsImport,
				fmt.Sprintf(
					"Baris %d: kode wajib diisi",
					rowIndex+1,
				),
			)
			continue
		}

		if row[6] != "" {
			qty, err := strconv.Atoi(strings.TrimSpace(row[6]))

			if err != nil {
				errorsImport = append(
					errorsImport,
					fmt.Sprintf(
						"baris %d : QTY harus berupa angka",
						rowIndex+1,
					),
				)
				continue
			}

			if qty < 0 {
				errorsImport = append(
					errorsImport,
					fmt.Sprintf(
						"baris %d : QTY tidak boleh negatif",
						rowIndex+1,
					),
				)
				continue
			}

			item.Qty =qty

		}	

			if row[8]!= "" {

				harga, err := strconv.ParseFloat(
					strings.ReplaceAll(
						strings.TrimSpace(row[8]),
						",",
						"",
					),
					64,
				)

				if err != nil {
					errorsImport = append(
						errorsImport,
						fmt.Sprintf(
							"baris %d : Harga + PPN harus berupa angka",
							rowIndex+1,
						),
					)
					continue
				}

				if harga < 0 {
					errorsImport = append(
						errorsImport,
						fmt.Sprintf(
							"baris%d : Harga + PPN tidak boleh negatif",
							rowIndex+1,
						),
					)
					continue
				}

				item.HargaPPN = harga
			}

			if row[14] != "" {

				jumlah, err := strconv.ParseFloat(
					strings.ReplaceAll(
						strings.TrimSpace(row[14]),
						",",
						"",
					),
					64,
				)

				if err != nil {
					errorsImport = append(
						errorsImport,
						fmt.Sprintf(
							"baris %d : Jumlah uang  masuk harus berupa angka",
							rowIndex+1,
						),
					)
					continue
				}

				if jumlah < 0 {
					errorsImport = append(
						errorsImport,
						fmt.Sprintf(
							"baris %d : JUmlah uang yang masuk tidak boleh negatif",
							rowIndex+1,
						),
					)
					continue
				}

				item.JumlahUangMasuk= jumlah
			}

			var dateErr error

			item.TglPesanan, dateErr = parseExcelDate(row[9])
			if dateErr != nil {
				errorsImport = append(
					errorsImport,
					fmt.Sprintf(
						"baris %d : TGL pesanan %s",
						rowIndex+1,
						dateErr.Error(),
					),
				)
				continue
			}

			item.TglBAST, dateErr = parseExcelDate(row[10])
			if dateErr != nil{
				errorsImport = append(
					errorsImport,
					fmt.Sprintf(
						"baris %d : TGL BAST %s",
						rowIndex+1,
						dateErr.Error(),
					),
				)
				continue
			}

			item.TglUangMasuk, dateErr = parseExcelDate(row[15])
			if dateErr != nil{
				errorsImport = append(
					errorsImport,
					fmt.Sprintf(
						"baris %d : TGL UANG MASUK %s",
						rowIndex+1,
						dateErr.Error(),
					),
				)
				continue
			}

			imported = append(imported, item)
		}

		if len(errorsImport) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":       false,
			"message":       "Import gagal karena terdapat data yang tidak valid",
			"total_valid":   len(imported),
			"total_error":   len(errorsImport),
			"errors":        errorsImport,
		})
		return
		}

		successCount := 0

		for i := range imported {

			if err := h.Service.Create(&imported[i]); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": fmt.Sprintf(
						"Gagal menyimpan data pada baris %d",
						i+2,
					),
					"error": err.Error(),
				})

				return
			}

			successCount++
		}

		c.JSON(http.StatusOK, gin.H{
			"success":      true,
			"message":      "Import Excel berhasil",
			"total_import": successCount,
		})
}


func parseExcelDate(value string) (models.Date, error) {

	value = strings.TrimSpace(value)

	if value == "" {
		return "", nil
	}

	
	if parsed, err := time.Parse("2006-01-02", value); err == nil {
		return models.Date(parsed.Format("2006-01-02")), nil
	}


	formats := []string{
		"02/01/2006",
		"02-01-2006",
		"01/02/2006",
		"2006/01/02",
	}

	for _, format := range formats {

		parsed, err := time.Parse(format, value)

		if err == nil {
			return models.Date(parsed.Format("2006-01-02")), nil
		}
	}

	return "", fmt.Errorf(
		"format tanggal tidak valid: %s, gunakan YYYY-MM-DD",
		value,
	)
}

func isEmptyExcelRow(row []string) bool {

	for _, value := range row {

		if strings.TrimSpace(value) != "" {
			return false
		}
	}

	return true
}


