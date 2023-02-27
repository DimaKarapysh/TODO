package domain

type User struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Pwd   string `json:"pwd"`
}

type UserForm struct {
	Login string `json:"login"binding:"required"`
	Pwd   string `json:"pwd"binding:"required"`
}

type Token struct {
	Token string `json:"token"binding:"required"`
}

type UserRepo interface {
	Insert(user *User) error
	FetchUserByName(userName string) (*User, error)
	FetchUserById(id int) (*User, error)
}

type UserIter interface {
	Register(userName, pwd string) error
	Login(userName, pwd string) (string, error)
	Verify(token string) (*User, error)
}
