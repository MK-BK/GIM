package models

import (
	"encoding/json"
)

const (
	CategorySystem = "SYSTEM"
	CategoryText   = "TEXT"
	CategoryImage  = "IMAGE"
)

type Message struct {
	Models
	Scope
	Category  string
	OwnerName string      `gorm:"-"`
	Spec      interface{} `gorm:"json"`
	Error     string
}

type MessageText struct {
	Content string
}

type MessageImage struct {
	Path string
}

func (m *Message) UnmarshalJSON(b []byte) error {
	var toMessage struct {
		Models
		Scope
		Category string
		Spec     json.RawMessage
		Error    string
	}

	if err := json.Unmarshal(b, &toMessage); err != nil {
		return err
	}

	m.Models = toMessage.Models
	m.Scope = toMessage.Scope
	m.Category = toMessage.Category
	m.Error = toMessage.Error

	switch toMessage.Category {
	case CategoryText:
		var spec MessageText
		if err := json.Unmarshal(toMessage.Spec, &spec); err != nil {
			return err
		}
		m.Spec = &spec
	case CategoryImage:
		var spec MessageImage
		if err := json.Unmarshal(toMessage.Spec, &spec); err != nil {
			return err
		}
		m.Spec = &spec
	}

	return nil
}

type MessageDistributionEvent struct {
	*Message
}

func (*MessageDistributionEvent) Type() string { return "message:distribution" }
