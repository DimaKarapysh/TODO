package interactors

import (
	"Portfolio_Nodes/domain"
	"sort"
)

type WinnerInteractor struct {
	winnerRepo domain.WinnerRepo
}

func NewWinnerInteractor(winnerRepo domain.WinnerRepo) *WinnerInteractor {
	return &WinnerInteractor{winnerRepo: winnerRepo}
}

func (i *WinnerInteractor) FindWinner() (domain.Winner, error) {
	all, err := i.winnerRepo.FetchAll()
	if err != nil {
		return domain.Winner{}, err
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].Points > all[j].Points
	})
	return all[0], nil
}
