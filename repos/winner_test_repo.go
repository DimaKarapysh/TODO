package repos

import (
	"Portfolio_Nodes/domain"
)

type WinnerTRepo struct {
}

func NewWinnerTRepo() *WinnerTRepo {
	return &WinnerTRepo{}
}

func (r *WinnerTRepo) FetchAll() ([]domain.Winner, error) {
	var winners []domain.Winner
	w1 := domain.Winner{
		Id:     4,
		Name:   "4535",
		Points: 435,
	}
	w2 := domain.Winner{
		Id:     234,
		Name:   "432",
		Points: 564,
	}
	winners = append(winners, w1)
	winners = append(winners, w2)
	return winners, nil
}
