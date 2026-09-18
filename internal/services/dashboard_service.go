package services

import (
	"time"

	"iot-backend/internal/repositories"
)

type DashboardService struct {
	Repository *repositories.DashboardRepository
}

func NewDashboardService(
	repository *repositories.DashboardRepository,
) *DashboardService {
	return &DashboardService{
		Repository: repository,
	}
}

func (s *DashboardService) GetDashboard(
	periode string,
) ([]repositories.DashboardResult, error) {

	now := time.Now()

	var startDate time.Time
	var endDate time.Time

	switch periode {

	case "day":

		startDate = time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			0,
			0,
			0,
			0,
			now.Location(),
		)

		endDate = startDate.AddDate(0, 0, 1)

	case "month":

		startDate = time.Date(
			now.Year(),
			now.Month(),
			1,
			0,
			0,
			0,
			0,
			now.Location(),
		)

		endDate = startDate.AddDate(0, 1, 0)

	case "year":

		startDate = time.Date(
			now.Year(),
			1,
			1,
			0,
			0,
			0,
			0,
			now.Location(),
		)

		endDate = startDate.AddDate(1, 0, 0)

	default:

		startDate = time.Date(
			now.Year(),
			now.Month(),
			1,
			0,
			0,
			0,
			0,
			now.Location(),
		)

		endDate = startDate.AddDate(0, 1, 0)

		periode = "month"
	}

	return s.Repository.GetDashboard(
		periode,
		startDate,
		endDate,
	)
}