package models

type Moment struct {
	Models
	DisplayName string `gorm:"-"`
	OwnerID     string
	Content     string    `gorm:"type:longtext"`
	Comments    []Comment `gorm:"-"`
}

type Moments interface {
	List(id string) ([]*Moment, error)
	Create(moment *Moment) error
	Get(id string) (*Moment, error)
	Delete(id string) error
}
