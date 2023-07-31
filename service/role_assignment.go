package service

import (
	"GIM/models"
	"GIM/pkg/config"
)

type RoleAssignmentManager struct{}

func NewRoleAssignmentManager() *RoleAssignmentManager {
	return &RoleAssignmentManager{}
}

func (*RoleAssignmentManager) CreateRoleAssignments(roles ...*models.RoleAssignment) error {
	return config.DB.Create(roles).Error
}

func (*RoleAssignmentManager) DeleteRoleAssignments(roles ...*models.RoleAssignment) error {
	return config.DB.Delete(roles).Error
}

func (*RoleAssignmentManager) GetRoleAssignments(role *models.RoleAssignment) ([]*models.RoleAssignment, error) {
	var roleAssignments []*models.RoleAssignment
	if err := config.DB.Where(role.Scope).Find(&roleAssignments).Error; err != nil {
		return nil, err
	}

	return roleAssignments, nil
}
