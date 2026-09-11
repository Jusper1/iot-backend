package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Date string

func (d *Date) UnmarshalJSON (data []byte) error {
	if string(data) == "null" {
		*d = ""
		return nil
	}

	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("tanggal harus berupa string dengan format YYYY-MM-DD")
	}

	if value == "" {
		*d = ""
		return nil
	}

	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return fmt.Errorf(
			"format tanggal tidak valid: %q, gunakan YYYY-MM-DD",
			value,
		)
	}

	formatted := parsed.Format("2006-01-02")

	if formatted != value {
		return fmt.Errorf(
			"format tanggal tidak valid: %q, gunakan YYYY-MM-DD",
			value,
		)
	}

	*d = Date(value)

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d == "" {
		return []byte("null"), nil
	}

	return json.Marshal(string(d))
}


func (d Date) Value() (driver.Value, error) {
	if d == "" {
		return nil, nil
	}

	if _, err := time.Parse("2006-01-02", string(d)); err != nil {
		return nil, fmt.Errorf("tanggal tidak valid: %s", d)
	}

	return string(d), nil
}

func (d *Date) Scan(value interface{}) error {
	if value == "" {
		*d = ""
		return nil
	}

	switch v := value.(type) {

	case time.Time:
		*d = Date(v.Format("2006-01-02"))
		return nil

	case []byte:
		date := string(v)

		if date == "" {
			*d = ""
			return nil
		}

		if _, err := time.Parse("2006-01-02", date); err != nil {
			return fmt.Errorf("tanggal database tidak valid: %s", date)
		}
		
		*d = Date(date)
		return nil

	case string:
		if v == "" {
			*d = ""
			return nil
		}

		if _, err := time.Parse("2006-01-02", v); err != nil {
			return fmt.Errorf("tanggal database tidak valid: %s", v)
		}

		*d = Date(v)
		return nil

	default:
		return fmt.Errorf("tidak dapat membaca tipe tanggal %T", value)
	}
}

type InaprocOrder struct {
	ID uint `json:"-" gorm:"primaryKey;autoIncrement"`

	Kode			  string  `json:"kode"`
	NamaPic			  string  `json:"nama_pic"`
	NIK				  string  `json:"nik"`
	Puskesmas		  string  `json:"puskesmas"`
	StatusOdoo		  string  `json:"status_odoo"`
	NoNPWP			  string  `json:"no_npwp"`
	Qty				  int	  `json:"qty"`
	Status            string  `json:"status"`
	HargaPPN          float64 `json:"harga_ppn"`
	TglPesanan        Date `json:"tgl_pesanan" gorm:"type:date"`
	TglBAST           Date `json:"tgl_bast" gorm:"type:date"`
	NoBAST            string  `json:"no_bast"`
	NoInvoiceKUT      string  `json:"no_invoice_kut"`
	NSFP              string  `json:"nsfp"`
	JumlahUangMasuk   float64 `json:"jumlah_uang_masuk"`
	TglUangMasuk      Date `json:"tgl_uang_masuk" gorm:"type:date"`
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