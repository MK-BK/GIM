package user

import (
	"GIM/models"
	"GIM/pkg/event"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func (m *UserManager) subscribe() {
	for evt := range event.Sub("request:update", "group:create", "group:delete") {
		if request, ok := evt.(*models.RequestUpdateEvent); ok {
			if request.Kind == models.ScopeUser {
				roles := []*models.RoleAssignment{
					{
						Scope: models.Scope{
							Kind:          models.ScopeUser,
							OwnerID:       request.OwnerID,
							DestinationID: request.DestinationID,
						},
					},
					{
						Scope: models.Scope{
							Kind:          models.ScopeUser,
							OwnerID:       request.DestinationID,
							DestinationID: request.OwnerID,
						},
					},
				}

				if err := models.GlobalEnvironment.RoleAssignments.CreateRoleAssignments(roles...); err != nil {
					log.Error(err)
				}
			} else {
				role := &models.RoleAssignment{
					Scope: models.Scope{
						Kind:          models.ScopeGroup,
						OwnerID:       request.OwnerID,
						DestinationID: request.GroupID,
					},
				}
				if err := models.GlobalEnvironment.RoleAssignments.CreateRoleAssignments(role); err != nil {
					log.Error(err)
				}
			}
		}

		if event, ok := evt.(*models.GroupCreateEvent); ok {
			roles := []*models.RoleAssignment{
				{
					Scope: models.Scope{
						Kind:    models.ScopeGroup,
						OwnerID: event.UserID,
						GroupID: event.GroupID,
					},
				},
			}

			for _, id := range event.UserIDs {
				roles = append(roles, &models.RoleAssignment{
					Scope: models.Scope{
						Kind:    models.ScopeGroup,
						OwnerID: id,
						GroupID: event.GroupID,
					},
				})
			}

			if err := models.GlobalEnvironment.RoleAssignments.CreateRoleAssignments(roles...); err != nil {
				log.Error(err)
			}
		}

		if event, ok := evt.(*models.GroupDeleteEvent); ok {
			role := &models.RoleAssignment{
				Scope: models.Scope{
					Kind:    models.ScopeGroup,
					GroupID: event.GroupID,
				},
			}
			if err := models.GlobalEnvironment.RoleAssignments.DeleteRoleAssignments(role); err != nil {
				log.Error(err)
			}
		}
	}
}
