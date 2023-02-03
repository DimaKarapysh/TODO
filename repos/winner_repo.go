package repos

import (
	"Portfolio_Nodes/domain"
	"gorm.io/gorm"
)

type WinnerRepo struct {
	db *gorm.DB
}

func NewWinnerRepo(db *gorm.DB) *WinnerRepo {
	return &WinnerRepo{db: db}
}

func (r *WinnerRepo) FetchAll() ([]domain.Winner, error) {
	var winners []domain.Winner
	result := r.db.Find(&winners)
	if result.Error != nil {
		return nil, result.Error
	}
	return winners, nil
}
