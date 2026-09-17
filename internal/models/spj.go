package models

import "time"

type SPJOrder struct {
	ID uint `json:"id" gorm:"primaryKey"`

	TglPrint 				Date `json:"tgl_print" gorm:"type:date"`
	TglParaf 				Date `json:"tgl_paraf" gorm:"type:date"`
	TglSign  				Date `json:"tgl_sign" gorm:"type:date"`
	TglKirim 				Date `json:"tgl_kirim" gorm:"type:date"`
	NamaDinkesPKM 			string `json:"nama_dinkes_pkm" gorm:"type:varchar(255);not null"`
	Kertas     	   			string `json:"kertas" gorm:"type:varchar(100)"`
	KebutuhanSPJ  			string `json:"kebutuhan_spj" gorm:"type:text"`
	JumlahRangkap 			uint   `json:"jumlah_rangkap"`
	JenisFile     			string `json:"jenis_file" gorm:"type:varchar(100)"`
	UPNomorTeleponAlamat 	string `json:"up_nomor_telepon_alamat" gorm:"type:text"`
	StatusPengiriman 		string `json:"status_pengiriman" gorm:"type:varchar(100)"`
	Keterangan       		string `json:"keterangan" gorm:"type:text"`
	PICPrint         		string `json:"pic_print" gorm:"type:varchar(255)"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (SPJOrder) TableName() string {
	return "spj_orders"
}