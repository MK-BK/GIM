package models

type Message struct {
	Models
	Scope
	OwnerName string `gorm:"-"`
	Content   string
	Error     string
}

type MessageDistributionEvent struct {
	*Message
}

func (*MessageDistributionEvent) Type() string { return "message:distribution" }
