package repos

import (
	"Portfolio_Nodes/domain"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Insert(user *domain.User) error {
	query := `INSERT INTO users (login,pwd) VALUES ($1,$2)`
	_, err := u.db.Exec(query, user.Login, user.Pwd)
	return err
}

func (u *UserRepo) FetchUserByName(userName string) (*domain.User, error) {
	var user = &domain.User{
		Login: userName,
	}
	query := `SELECT id,login,pwd FROM users WHERE login=$1`
	err := u.db.QueryRow(query, userName).Scan(&user.Id, &user.Login, &user.Pwd)
	switch err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (u *UserRepo) FetchUserById(id int) (*domain.User, error) {
	var user = &domain.User{
		Id: id,
	}
	query := `SELECT id,login,pwd FROM users WHERE id=$1`
	err := u.db.QueryRow(query, id).Scan(&user.Id, &user.Login, &user.Pwd)
	switch err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}
