package service

import (
	"errors"

	"GIM/models"
	"GIM/pkg/cache"
	"GIM/pkg/config"
	"GIM/pkg/event"
)

type groupManager struct{}

var _ models.Groups = (*groupManager)(nil)

func NewGroupManager() *groupManager {
	return &groupManager{}
}

func (*groupManager) ListGroups(id string) ([]*models.Group, error) {
	roles, err := models.GlobalEnvironment.RoleAssignments.GetRoleAssignments(&models.RoleAssignment{
		Scope: models.Scope{
			Kind:    models.ScopeGroup,
			OwnerID: id,
		},
	})
	if err != nil {
		return nil, err
	}

	releations := make([]string, 0)
	for _, role := range roles {
		releations = append(releations, role.GroupID)
	}

	groups := make([]*models.Group, 0)
	if len(releations) > 0 {
		if err := config.DB.Where("id IN ( ? )", releations).Find(&groups).Error; err != nil {
			return nil, err
		}
	}

	return groups, nil
}

func (m *groupManager) CreateGroup(userIDs []string, group *models.Group) (*models.Group, error) {
	if err := config.DB.Save(&group).Error; err != nil {
		return nil, err
	}

	event.Pub(&models.GroupCreateEvent{
		UserID:  group.ManagerID,
		GroupID: group.ID,
		UserIDs: userIDs,
	})

	return group, nil
}

func (*groupManager) UpdateGroup(id string, patch *models.GroupPatch) error {
	var group models.Group
	if err := config.DB.Where("id = ?", id).First(&group).Error; err != nil {
		return err
	}

	return config.DB.Model(&group).Updates(map[string]interface{}{
		"display_name": patch.Name,
		"announcement": patch.Announcement,
		"verify":       patch.Verify,
	}).Error
}

func (m *groupManager) GetGroup(id string) (*models.Group, error) {
	group, exist := cache.GetGroup(id)
	if exist {
		return group, nil
	}

	if err := config.DB.Where("id = ?", id).Find(&group).Error; err != nil {
		return nil, err
	}

	roles, err := models.GlobalEnvironment.RoleAssignments.GetRoleAssignments(&models.RoleAssignment{
		Scope: models.Scope{
			Kind:    models.ScopeGroup,
			GroupID: id,
		},
	})
	if err != nil {
		log.Error(err)
	}

	if len(roles) > 0 {
		releations := make([]string, 0)
		for _, role := range roles {
			releations = append(releations, role.OwnerID)
		}

		if err := config.DB.Find(&group.Users, releations).Error; err != nil {
			return nil, err
		}
	}

	return group, nil
}

func (*groupManager) DeleteGroup(userID, id string) error {
	var group models.Group
	if err := config.DB.Find(&group, "id = ?", id).Error; err != nil {
		return err
	}

	if userID != group.ManagerID {
		return errors.New("not allow to delete room")
	}

	if err := config.DB.Delete(&models.Group{}, "id = ?", id).Error; err != nil {
		return err
	}

	event.Pub(&models.GroupDeleteEvent{
		GroupID: id,
	})

	return nil
}

func (m *groupManager) JoinGroup(userID, groupID string, userIDs []string) error {
	group, err := m.GetGroup(groupID)
	if err != nil {
		return err
	}

	if userID == group.ManagerID || !group.Verify {
		roles := make([]*models.RoleAssignment, 0)
		for _, id := range userIDs {
			roles = append(roles, &models.RoleAssignment{
				Scope: models.Scope{
					Kind:    models.ScopeGroup,
					OwnerID: id,
					GroupID: groupID,
				},
			})
		}

		if err := models.GlobalEnvironment.RoleAssignments.CreateRoleAssignments(roles...); err != nil {
			return err
		}
	} else {
		for _, userID := range userIDs {
			request := &models.Request{
				Scope: models.Scope{
					Kind:          models.ScopeGroup,
					OwnerID:       userID,
					DestinationID: group.ManagerID,
					GroupID:       groupID,
				},
			}

			if err := models.GlobalEnvironment.Requests.CreateRequest(request); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *groupManager) LeaveGroup(userID, groupID, deleteUserID string) error {
	group, err := m.GetGroup(groupID)
	if err != nil {
		return err
	}

	if userID == deleteUserID || userID == group.ManagerID {
		role := &models.RoleAssignment{
			Scope: models.Scope{
				Kind:    models.ScopeGroup,
				OwnerID: deleteUserID,
				GroupID: groupID,
			},
		}
		return models.GlobalEnvironment.RoleAssignments.DeleteRoleAssignments(role)
	}

	return errors.New("not allow to delete user from Group")
}
