package models

import "time"


type ManualOrder struct{
	ID uint64`gorm:"PrimaryKey" json:"id"`

	Kode					string `json:"kode"`
	DinkesPuskesmas 		string `json:"dinkes_puskesmas"`
	Email           		string `json:"email"`
	StatusOdoo				string `json:"status_odoo"`
	Status    				string `json:"status"`
	Qty                 	uint   `json:"qty"`
	WaktuPengirimanBulan	*uint  `json:"waktu_pengiriman_bulan"`
	NamaPIC					string `json:"nama_pic"`
	NIK   					string `json:"nik"`
	NPWP   					string `json:"npwp"`
	NoPO 					string `json:"no_po"`
	TglPO					Date   `json:"tgl_po"`
	NoBAST					string `json:"no_bast"`
	TglBAST					Date   `json:"tgl_bast"`
	HargaPPN				float64 `json:"harga_ppn"`
	TglUangMasuk			Date 	`json:"tgl_uang_masuk"`
	JumlahUangMasuk  		float64 `json:"jumlah_uang_masuk"`
	PeriodeLangganan		string 	`json:"periode_langganan"`
	NomorInvoice			string 	`json:"nomor_invoice"`
	Alamat					string 	`json:"alamat"`
	KotaKab 				string 	`json:"kota_kab"`
	
	CreatedAt 				time.Time  `json:"created_at"`
	UpdatedAt 				time.Time  `json:"updated_at"`
}

func (ManualOrder) TableName() string {
	return "pemesanan_manual" 
}