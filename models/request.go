package models

type Request struct {
	Models
	Scope
	Description string
	OwnerName   string `gorm:"-"`
	Status      bool
}

type Requests interface {
	ListRequests(userID string) ([]*Request, error)
	CreateRequest(request *Request) error
	GetRequest(id string) (*Request, error)
	UpdateRequest(userID, id string) error
	DeleteRequest(id string) error
	GetDescription(request *Request) string
}

type RequestUpdateEvent struct {
	*Request
}

func (*RequestUpdateEvent) Type() string { return "request:update" }
