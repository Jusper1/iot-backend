package repositories

import (
	"time"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{
		DB: db,
	}
}

type DashboardResult struct {
	Periode         string `json:"periode"`
	Inaproc         int64  `json:"inaproc"`
	Manual          int64  `json:"manual"`
	Alat            int64  `json:"alat"`
	InaprocDibayar  int64  `json:"inaproc_dibayar"`
	ManualDibayar   int64  `json:"manual_dibayar"`
	AlatDibayar     int64  `json:"alat_dibayar"`
}

func (r *DashboardRepository) GetDashboard(
	periode string,
	startDate time.Time,
	endDate time.Time,
) ([]DashboardResult, error) {

	var result []DashboardResult

	var dateFormat string

	switch periode {
	case "day":
		dateFormat = "%Y-%m-%d"

	case "month":
		dateFormat = "%Y-%m"

	case "year":
		dateFormat = "%Y"

	default:
		dateFormat = "%Y-%m"
	}

	query := `
		SELECT
			periode,

			SUM(inaproc) AS inaproc,
			SUM(manual) AS manual,
			SUM(alat) AS alat,

			SUM(inaproc_dibayar) AS inaproc_dibayar,
			SUM(manual_dibayar) AS manual_dibayar,
			SUM(alat_dibayar) AS alat_dibayar

		FROM (


			SELECT
				DATE_FORMAT(tgl_pesanan, ?) AS periode,

				1 AS inaproc,
				0 AS manual,
				0 AS alat,

				CASE
					WHEN status = 'paid'
					THEN 1
					ELSE 0
				END AS inaproc_dibayar,

				0 AS manual_dibayar,
				0 AS alat_dibayar

			FROM inaproc_orders

			WHERE tgl_pesanan >= ?
			  AND tgl_pesanan < ?


			UNION ALL


			SELECT
				DATE_FORMAT(tgl_po, ?) AS periode,

				0 AS inaproc,
				1 AS manual,
				0 AS alat,

				0 AS inaproc_dibayar,

				CASE
					WHEN status = 'paid'
					THEN 1
					ELSE 0
				END AS manual_dibayar,

				0 AS alat_dibayar

			FROM manual_orders

			WHERE tgl_po >= ?
			  AND tgl_po < ?


			UNION ALL


			SELECT
				DATE_FORMAT(tgl_po, ?) AS periode,

				0 AS inaproc,
				0 AS manual,
				1 AS alat,

				0 AS inaproc_dibayar,
				0 AS manual_dibayar,

				CASE
					WHEN LOWER(status_pembayaran) IN (
						'paid',
						'lunas',
						'sudah dibayar',
						'sudah bayar'
					)
					THEN 1
					ELSE 0
				END AS alat_dibayar

			FROM pembelian_alat

			WHERE tgl_po >= ?
			  AND tgl_po < ?

		) AS data

		GROUP BY periode

		ORDER BY periode ASC
	`

	err := r.DB.Raw(
		query,

		dateFormat,
		startDate,
		endDate,

		dateFormat,
		startDate,
		endDate,

		dateFormat,
		startDate,
		endDate,
	).Scan(&result).Error

	return result, err
}