package domain

type Winner struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type WinnerRepo interface {
	FetchAll() ([]Winner, error)
}

type WinnerInteractor interface {
	FindWinner() (Winner, error)
}
