package user

import (
	"fmt"

	"GIM/models"
	"GIM/pkg/config"
)

func (m *UserManager) AddFriend(userID string, options *models.UserAddOption) error {
	if options.Description == "" {
		user, err := m.GetUser(userID)
		if err != nil {
			return err
		}
		options.Description = fmt.Sprintf("From user: %s", user.Name)
	}

	request := &models.Request{
		Description: options.Description,
		Scope: models.Scope{
			Kind:          models.ScopeUser,
			OwnerID:       userID,
			DestinationID: options.ID,
		},
	}

	return models.GlobalEnvironment.Requests.CreateRequest(request)
}

func (m *UserManager) ListFriends(id string) ([]*models.User, error) {
	var users []*models.User

	err := config.DB.Model(&models.User{}).
		Joins("left join role_assignments on role_assignments.destination_id = users.id").
		Where("role_assignments.kind = ? AND role_assignments.owner_id = ?", models.ScopeUser, id).
		Scan(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserManager) DeleteFriend(userID, id string) error {
	roles := []*models.RoleAssignment{
		{
			Scope: models.Scope{
				Kind:          models.ScopeUser,
				OwnerID:       id,
				DestinationID: userID,
			},
		},
		{
			Scope: models.Scope{
				Kind:          models.ScopeUser,
				OwnerID:       userID,
				DestinationID: id,
			},
		}}

	return models.GlobalEnvironment.RoleAssignments.DeleteRoleAssignments(roles...)
}
