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
	Error     string      `gorm:"json"`
}

type MessageText struct {
	Content string
}

type MessageImage struct {
	Path string
}

func (message *Message) UnmarshalJSON(b []byte) error {
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

	message.Models = toMessage.Models
	message.Scope = toMessage.Scope
	message.Category = toMessage.Category
	message.Error = toMessage.Error

	switch toMessage.Category {
	case CategoryText:
		var spec MessageText
		if err := json.Unmarshal(toMessage.Spec, &spec); err != nil {
			return err
		}
		message.Spec = &spec
	case CategoryImage:
		var spec MessageImage
		if err := json.Unmarshal(toMessage.Spec, &spec); err != nil {
			return err
		}
		message.Spec = &spec
	}

	return nil
}

type MessageDistributionEvent struct {
	*Message
}

func (*MessageDistributionEvent) Type() string { return "message:distribution" }

func (message *MessageDistributionEvent) GetUserID() string {
	if message.Kind == ScopeUser {
		return message.DestinationID
	}

	return message.GroupID
}
