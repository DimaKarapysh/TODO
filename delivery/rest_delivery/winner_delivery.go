package rest_delivery

import (
	"Portfolio_Nodes/domain"
	"github.com/gin-gonic/gin"
)

type WinnerDelivery struct {
	winnerInteractor domain.WinnerInteractor
}

func NewWinnerDelivery(winnerInteractor domain.WinnerInteractor) *WinnerDelivery {
	return &WinnerDelivery{winnerInteractor: winnerInteractor}
}

func (s *WinnerDelivery) Route(r *gin.RouterGroup) {
	r.POST("/get", s.CalculateMax)
}

func (s *WinnerDelivery) CalculateMax(context *gin.Context) {
	winner, err := s.winnerInteractor.FindWinner()
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(200, MakeSuccessWithData(winner))
}
