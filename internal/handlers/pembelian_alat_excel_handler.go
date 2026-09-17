package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type PembelianAlatExcelHandler struct {
	Service *services.PembelianAlatService
}

func NewPembelianAlatExcelHandler(
	service *services.PembelianAlatService,
) *PembelianAlatExcelHandler {
	return &PembelianAlatExcelHandler{
		Service: service,
	}
}

func (h *PembelianAlatExcelHandler) Export(c *gin.Context) {

	data, err := h.Service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal mengambil data",
			"error":   err.Error(),
		})
		return
	}

	file := excelize.NewFile()

	sheet := "Pembelian Alat"

	index, err := file.NewSheet(sheet)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal membuat sheet Excel",
			"error":   err.Error(),
		})
		return
	}

	file.DeleteSheet("Sheet1")
	file.SetActiveSheet(index)

	headers := []string{
		"nomor",
		"Instansi/Perusahaan",
		"email",
		"Status",
		"No PO KUT",
		"PO ke Tsel",
		"Invoice",
		"nama PIC",
		"Jabatan",
		"nik",
		"npwp",
		"no PO",
		"tgl PO",
		"waktu tiba",
		"selisih waktu",
		"no BAST",
		"tgl BAST",
		"no invoice",
		"keterangan",
		"tipe timbangan",
		"Quantity",
		"harga produk timbangan",
		"harga ongkir KUT",
		"harga produk dikali qty",
		"ppn produk 11%",
		"harga ongkir dikali qty",
		"ppn ongkir 11%",
		"total harga + ongkir",
		"total harga jual",
		"uang masuk",
		"tgl uang masuk",
		"kode bayar",
		"no invoice inaproc",
		"status pembayaran",
		"wilayah pengiriman",
		"berat (KG)",
		"harga produk (reseller)",
		"harga produk dikali qty (reseller)",
		"harga ongkir(reseller)",
		"harga ongkir dikali qty(reseller)",
		"ppn ongkir 11% (reseller)",
		"ppn produk 11% (reseller)",
		"harga produk + ongkir (reseller)",
		"total harga reseller + ppn",
		"alamat",
		"Status Pengiriman",
		"Nama Ekspedisi",
		"No Resi",
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)

		if err := file.SetCellValue(
			sheet,
			cell,
			header,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "gagal menulis header Excel",
				"error":   err.Error(),
			})
			return
		}
	}

	for rowIndex, item := range data {

		row := rowIndex + 2

		values := []interface{}{
			item.Nomor,
			item.InstansiPerusahaan,
			item.Email,
			item.Status,
			item.NoPOKUT,
			item.POKeTsel,
			item.Invoice,
			item.NamaPIC,
			item.Jabatan,
			item.NIK,
			item.NPWP,
			item.NoPO,
			item.TglPO,
			item.WaktuTiba,
			item.SelisihWaktu,
			item.NoBAST,
			item.TglBAST,
			item.NoInvoice,
			item.Keterangan,
			item.TipeTimbangan,
			item.Quantity,
			item.HargaProdukTimbangan,
			item.HargaOngkirKUT,
			item.HargaProdukDikaliQty,
			item.PPNProduk11,
			item.HargaOngkirDikaliQty,
			item.PPNOngkir11,
			item.TotalHargaOngkir,
			item.TotalHargaJual,
			item.UangMasuk,
			item.TglUangMasuk,
			item.KodeBayar,
			item.NoInvoiceInaproc,
			item.StatusPembayaran,
			item.WilayahPengiriman,
			item.BeratKG,
			item.HargaProdukReseller,
			item.HargaProdukDikaliQtyReseller,
			item.HargaOngkirReseller,
			item.HargaOngkirDikaliQtyReseller,
			item.PPNOngkir11Reseller,
			item.PPNProduk11Reseller,
			item.HargaProdukOngkirReseller,
			item.TotalHargaResellerPPN,
			item.Alamat,
			item.StatusPengiriman,
			item.NamaEkspedisi,
			item.NoResi,
		}

		for colIndex, value := range values {

			cell, _ :=
				excelize.CoordinatesToCellName(
					colIndex+1,
					row,
				)

			if err := file.SetCellValue(
				sheet,
				cell,
				value,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "gagal menulis data Excel",
					"error":   err.Error(),
				})
				return
			}
		}
	}

	c.Header(
		"Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	)

	c.Header(
		"Content-Disposition",
		`attachment; filename="pembelian_alat.xlsx"`,
	)

	if err := file.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal mengirim file Excel",
			"error":   err.Error(),
		})
		return
	}
}

func (h *PembelianAlatExcelHandler) Import(c *gin.Context) {

	fileHeader, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "file Excel wajib diupload",
			"error":   err.Error(),
		})
		return
	}

	file, err := fileHeader.Open()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "file Excel tidak dapat dibuka",
			"error":   err.Error(),
		})
		return
	}

	defer file.Close()

	excel, err := excelize.OpenReader(file)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "file Excel tidak valid",
			"error":   err.Error(),
		})
		return
	}

	defer excel.Close()

	sheet := excel.GetSheetName(0)

	if sheet == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "sheet Excel tidak ditemukan",
		})
		return
	}

	rows, err := excel.GetRows(sheet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal membaca Excel",
			"error":   err.Error(),
		})
		return
	}

	if len(rows) <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Excel tidak memiliki data",
		})
		return
	}

	imported := 0
	errorsImport := []string{}

	for i, row := range rows {


		if i == 0 {
			continue
		}

		if len(row) == 0 {
			continue
		}

		data := models.PembelianAlat{

			Nomor: int(getExcelInt(row, 0)),

			InstansiPerusahaan:
				getExcelString(row, 1),

			Email:
				getExcelString(row, 2),

			Status:
				getExcelString(row, 3),

			NoPOKUT:
				getExcelString(row, 4),

			POKeTsel:
				getExcelString(row, 5),

			Invoice:
				getExcelString(row, 6),

			NamaPIC:
				getExcelString(row, 7),

			Jabatan:
				getExcelString(row, 8),

			NIK:
				getExcelString(row, 9),

			NPWP:
				getExcelString(row, 10),

			NoPO:
				getExcelString(row, 11),

			TglPO:
				models.Date(getExcelString(row, 12)),

			WaktuTiba:
				models.Date(getExcelString(row, 13)),

			SelisihWaktu:
				int(getExcelInt(row, 14)),

			NoBAST:
				getExcelString(row, 15),

			TglBAST:
				models.Date(getExcelString(row, 16)),

			NoInvoice:
				getExcelString(row, 17),

			Keterangan:
				getExcelString(row, 18),

			TipeTimbangan:
				getExcelString(row, 19),

			Quantity:
				uint(getExcelInt(row, 20)),

			HargaProdukTimbangan:
				getExcelFloat(row, 21),

			HargaOngkirKUT:
				getExcelFloat(row, 22),

			HargaProdukDikaliQty:
				getExcelFloat(row, 23),

			PPNProduk11:
				getExcelFloat(row, 24),

			HargaOngkirDikaliQty:
				getExcelFloat(row, 25),

			PPNOngkir11:
				getExcelFloat(row, 26),

			TotalHargaOngkir:
				getExcelFloat(row, 27),

			TotalHargaJual:
				getExcelFloat(row, 28),

			UangMasuk:
				getExcelFloat(row, 29),

			TglUangMasuk:
				models.Date(getExcelString(row, 30)),

			KodeBayar:
				getExcelString(row, 31),

			NoInvoiceInaproc:
				getExcelString(row, 32),

			StatusPembayaran:
				getExcelString(row, 33),

			WilayahPengiriman:
				getExcelString(row, 34),

			BeratKG:
				getExcelFloat(row, 35),

			HargaProdukReseller:
				getExcelFloat(row, 36),

			HargaProdukDikaliQtyReseller:
				getExcelFloat(row, 37),

			HargaOngkirReseller:
				getExcelFloat(row, 38),

			HargaOngkirDikaliQtyReseller:
				getExcelFloat(row, 39),

			PPNOngkir11Reseller:
				getExcelFloat(row, 40),

			PPNProduk11Reseller:
				getExcelFloat(row, 41),

			HargaProdukOngkirReseller:
				getExcelFloat(row, 42),

			TotalHargaResellerPPN:
				getExcelFloat(row, 43),

			Alamat:
				getExcelString(row, 44),

			StatusPengiriman:
				getExcelString(row, 45),

			NamaEkspedisi:
				getExcelString(row, 46),

			NoResi:
				getExcelString(row, 47),
		}

		if err := h.Service.Create(&data); err != nil {
			errorsImport = append(
				errorsImport,
				fmt.Sprintf(
					"baris %d: %s",
					i+1,
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