package service

import (
	"GIM/models"
)

type RoleAssignmentManager struct{}

func NewRoleAssignmentManager() *RoleAssignmentManager {
	return &RoleAssignmentManager{}
}

func (*RoleAssignmentManager) CreateRoleAssignments(roles ...*models.RoleAssignment) error {
	return db.Create(roles).Error
}

func (*RoleAssignmentManager) DeleteRoleAssignments(roles ...*models.RoleAssignment) error {
	return db.Delete(roles).Error
}

func (*RoleAssignmentManager) GetRoleAssignments(role *models.RoleAssignment) ([]*models.RoleAssignment, error) {
	var roleAssignments []*models.RoleAssignment
	if err := db.Model(role).First(&roleAssignments).Error; err != nil {
		return nil, err
	}

	return roleAssignments, nil
}
