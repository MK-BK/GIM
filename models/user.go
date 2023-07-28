package models

type User struct {
	Models
	Name     string `sql:"index"`
	Password string
	Email    string `sql:"index"`
}

type UserPatch struct {
	Name     string
	Password string
	Email    string
}

type UserListOptions struct {
	Search string
}

type UserAddOption struct {
	ID          string
	Description string
}

type Users interface {
	Login(user *User) (string, error)
	Logout(userID string) error
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
	ListUsers(option *UserListOptions) ([]*User, error)
	UpdateUser(id string, patch *UserPatch) error
	ListFriends(userID string) ([]*User, error)
	AddFriend(userID string, option *UserAddOption) error
	DeleteFriend(userID, id string) error
}
