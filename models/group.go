package models

type Group struct {
	Models
	DisplayName  string
	Announcement string  // 系统公告
	ManagerID    string  // 管理员ID
	Verify       bool    // 开启审核入群， 需要管理员进行审核
	Users        []*User `gorm:"-"`
}

type Groups interface {
	ListGroups(id string) ([]*Group, error)
	CreateGroup(users []string, group *Group) (*Group, error)
	UpdateGroup(id string, patch *GroupPatch) error
	GetGroup(id string) (*Group, error)
	DeleteGroup(userID, id string) error

	JoinGroup(userID, groupID string, userIDs []string) error
	LeaveGroup(userID, groupID, users string) error
}

type GroupPatch struct {
	Name         string `json:"name"`
	Announcement string `json:"announcement"`
	Verify       bool   `json:"verify"`
}

// Group Event
type GroupDeleteEvent struct {
	GroupID string
}

func (*GroupDeleteEvent) Type() string { return "group:delete" }

type GroupCreateEvent struct {
	UserID  string
	GroupID string
	UserIDs []string
}

func (*GroupCreateEvent) Type() string { return "group:create" }
