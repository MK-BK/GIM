package models

type RoleAssignment struct {
	Models
	Scope
}

type RoleAssignments interface {
	CreateRoleAssignments(roles ...*RoleAssignment) error
	DeleteRoleAssignments(roles ...*RoleAssignment) error
	GetRoleAssignments(role *RoleAssignment) ([]*RoleAssignment, error)
}
