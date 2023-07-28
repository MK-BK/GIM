package models

type Session struct {
	Models
	Scope
	DisplayName   string
	LatestMessage string
	Enabled       bool `gorm:"-"`
}

type Sessions interface {
	List(userID string) ([]*Session, error)
	Create(session *Session) (*Session, error)
	Get(id string) (*Session, error)
	Delete(id string) error
}
