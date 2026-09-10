package models

import "time"

type InaprocOrder struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Kode			  string  `json:"kode"`
	NamaPic			  string  `json:"nama_pic"`
	NIK				  string  `json:"nik"`
	Puskesmas		  string  `json:"puskesmas"`
	StatusOdoo		  string  `json:"status_odoo"`
	NoNPWP			  string  `json:"no_npwp"`
	Qty				  int	  `json:"qty"`
	Status            string  `json:"status"`
	HargaPPN          float64 `json:"harga_ppn"`
	TglPesanan        *time.Time `json:"tgl_pesanan"`
	TglBAST           *time.Time `json:"tgl_bast"`
	NoBAST            string  `json:"no_bast"`
	NoInvoiceKUT      string  `json:"no_invoice_kut"`
	NSFP              string  `json:"nsfp"`
	JumlahUangMasuk   float64 `json:"jumlah_uang_masuk"`
	TglUangMasuk      *time.Time `json:"tgl_uang_masuk"`
	Rekening          string  `json:"rekening"`
	NoInvoiceInaproc  string  `json:"no_invoice_inaproc"`
	KodeBayar         string  `json:"kode_bayar"`
	Keterangan        string  `json:"keterangan"`
	Alamat            string  `json:"alamat"`
	Provinsi          string  `json:"provinsi"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (InaprocOrder) TableName() string {
	return "inaproc_orders"
}