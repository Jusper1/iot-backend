package models

import "time"

type PembelianAlat struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Nomor              int    `json:"nomor" gorm:"column:nomor"`
	InstansiPerusahaan string `json:"instansi_perusahaan" gorm:"column:instansi_perusahaan"`
	Email              string `json:"email" gorm:"column:email"`
	Status             string `json:"status" gorm:"column:status"`

	NoPOKUT  string `json:"no_po_kut" gorm:"column:no_po_kut"`
	POKeTsel string `json:"po_ke_tsel" gorm:"column:po_ke_tsel"`
	Invoice  string `json:"invoice" gorm:"column:invoice"`

	NamaPIC string `json:"nama_pic" gorm:"column:nama_pic"`
	Jabatan string `json:"jabatan" gorm:"column:jabatan"`
	NIK     string `json:"nik" gorm:"column:nik"`
	NPWP    string `json:"npwp" gorm:"column:npwp"`

	NoPO        string `json:"no_po" gorm:"column:no_po"`
	TglPO       Date   `json:"tgl_po" gorm:"column:tgl_po;type:date"`
	WaktuTiba   Date   `json:"waktu_tiba" gorm:"column:waktu_tiba;type:date"`
	SelisihWaktu int   `json:"selisih_waktu" gorm:"column:selisih_waktu"`

	NoBAST     string `json:"no_bast" gorm:"column:no_bast"`
	TglBAST    Date   `json:"tgl_bast" gorm:"column:tgl_bast;type:date"`
	NoInvoice  string `json:"no_invoice" gorm:"column:no_invoice"`
	Keterangan string `json:"keterangan" gorm:"column:keterangan"`

	TipeTimbangan string `json:"tipe_timbangan" gorm:"column:tipe_timbangan"`
	Quantity      uint   `json:"quantity" gorm:"column:quantity"`

	HargaProdukTimbangan float64 `json:"harga_produk_timbangan" gorm:"column:harga_produk_timbangan"`
	HargaOngkirKUT       float64 `json:"harga_ongkir_kut" gorm:"column:harga_ongkir_kut"`
	HargaProdukDikaliQty float64 `json:"harga_produk_dikali_qty" gorm:"column:harga_produk_dikali_qty"`
	PPNProduk11          float64 `json:"ppn_produk_11" gorm:"column:ppn_produk_11"`
	HargaOngkirDikaliQty float64 `json:"harga_ongkir_dikali_qty" gorm:"column:harga_ongkir_dikali_qty"`
	PPNOngkir11          float64 `json:"ppn_ongkir_11" gorm:"column:ppn_ongkir_11"`

	TotalHargaOngkir float64 `json:"total_harga_ongkir" gorm:"column:total_harga_ongkir"`
	TotalHargaJual   float64 `json:"total_harga_jual" gorm:"column:total_harga_jual"`

	UangMasuk    float64 `json:"uang_masuk" gorm:"column:uang_masuk"`
	TglUangMasuk Date    `json:"tgl_uang_masuk" gorm:"column:tgl_uang_masuk;type:date"`

	KodeBayar        string `json:"kode_bayar" gorm:"column:kode_bayar"`
	NoInvoiceInaproc string `json:"no_invoice_inaproc" gorm:"column:no_invoice_inaproc"`
	StatusPembayaran string `json:"status_pembayaran" gorm:"column:status_pembayaran"`

	WilayahPengiriman string  `json:"wilayah_pengiriman" gorm:"column:wilayah_pengiriman"`
	BeratKG           float64 `json:"berat_kg" gorm:"column:berat_kg"`

	HargaProdukReseller          float64 `json:"harga_produk_reseller" gorm:"column:harga_produk_reseller"`
	HargaProdukDikaliQtyReseller float64 `json:"harga_produk_dikali_qty_reseller" gorm:"column:harga_produk_dikali_qty_reseller"`
	HargaOngkirReseller          float64 `json:"harga_ongkir_reseller" gorm:"column:harga_ongkir_reseller"`
	HargaOngkirDikaliQtyReseller float64 `json:"harga_ongkir_dikali_qty_reseller" gorm:"column:harga_ongkir_dikali_qty_reseller"`
	PPNOngkir11Reseller          float64 `json:"ppn_ongkir_11_reseller" gorm:"column:ppn_ongkir_11_reseller"`
	PPNProduk11Reseller          float64 `json:"ppn_produk_11_reseller" gorm:"column:ppn_produk_11_reseller"`
	HargaProdukOngkirReseller    float64 `json:"harga_produk_ongkir_reseller" gorm:"column:harga_produk_ongkir_reseller"`
	TotalHargaResellerPPN        float64 `json:"total_harga_reseller_ppn" gorm:"column:total_harga_reseller_ppn"`

	Alamat           string `json:"alamat" gorm:"column:alamat"`
	StatusPengiriman string `json:"status_pengiriman" gorm:"column:status_pengiriman"`
	NamaEkspedisi    string `json:"nama_ekspedisi" gorm:"column:nama_ekspedisi"`
	NoResi           string `json:"no_resi" gorm:"column:no_resi"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (PembelianAlat) TableName() string {
	return "pembelian_alat"
}