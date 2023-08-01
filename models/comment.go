package models

type Comment struct {
	Models
	OwnerID  string
	MomentID string
	Content  string `gorm:"type:longtext"`
}

type Comments interface {
	List(id string) ([]*Comment, error)
	Create(comment *Comment) error
	Delete(id string) error
}
