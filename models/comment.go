package models

type Comment struct {
	Models
	OwnerID  string
	MomentID string
	ParentID string
	Content  string `gorm:"type:longtext"`
}

type TreeComment struct {
	*Comment
	Children []*TreeComment
}

type Comments interface {
	List(id string) ([]*Comment, error)
	Create(comment *Comment) error
	Delete(id string) error
}
