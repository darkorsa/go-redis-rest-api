package domain

func NewUser(username string, password string) *User {
	return &User{
		Username: username,
		password: password,
	}
}

type User struct {
	Username string `json:"username"`
	password string
}

func (u *User) GetPassword() string {
	return u.password
}
