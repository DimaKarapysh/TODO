package interactors

import (
	"Portfolio_Nodes/domain"
	"Portfolio_Nodes/tools"
	"errors"
)

type UserIter struct {
	domain.UserRepo
}

func NewUserIter(userRepo domain.UserRepo) *UserIter {
	return &UserIter{UserRepo: userRepo}
}

func (c *UserIter) Register(userName, pwd string) error {
	if userName == "" || pwd == "" {
		return errors.New("validation error")
	}
	existsUser, err := c.UserRepo.FetchUserByName(userName)
	if err != nil {
		return err
	}
	if existsUser != nil {
		return errors.New("already exist")
	}
	pwd, err = tools.GenerateHashPassword(pwd)
	if err != nil {
		return err
	}
	user := &domain.User{
		Login: userName,
		Pwd:   pwd,
	}
	err = c.UserRepo.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserIter) Login(userName, pwd string) (string, error) {
	user, err := c.UserRepo.FetchUserByName(userName)
	if err != nil {
		return "", err
	}
	if user == nil {
		err = errors.New("not found login")
		return "", err
	}
	if tools.CheckPasswordHash(user.Pwd, pwd) {
		return "", errors.New("not true password")
	}
	token, err := tools.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *UserIter) Verify(token string) (*domain.User, error) {

	ok, id, err := tools.ExtractData(token)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic(errors.New("not ok"))
	}
	user, err := c.UserRepo.FetchUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
