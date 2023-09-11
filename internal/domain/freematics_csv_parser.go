package domain

import "github.com/tomhobson/freematics2prom/pkg/models"

type FreematicsCSVParser interface {
	ParseCSV(csvParser string) ([]models.FreematicsData, error)
}
