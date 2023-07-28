package service

import (
	"strings"

	"GIM/models"
	"GIM/pkg/event"
)

type SessionManager struct{}

var _ models.Sessions = (*SessionManager)(nil)

func NewSessionManager() *SessionManager {
	s := &SessionManager{}
	go s.subscribe()
	return s
}

func (s *SessionManager) List(userID string) ([]*models.Session, error) {
	var sessions []*models.Session
	if err := db.Where("owner_id = ?", userID).Order("updated_at DESC").Find(&sessions).Error; err != nil {
		return nil, err
	}

	for _, session := range sessions {
		if session.Kind == models.ScopeUser {
			user, err := models.GlobalEnvironment.Users.GetUser(session.DestinationID)
			if err != nil {
				log.Error(err)
				continue
			}
			session.DisplayName = user.Name
		}

		if session.Kind == models.ScopeGroup {
			group, err := models.GlobalEnvironment.Groups.GetGroup(session.GroupID)
			if err != nil {
				log.Error(err)
				continue
			}

			if group.DisplayName != "" {
				session.DisplayName = group.DisplayName
			} else {
				names := make([]string, 0)
				if session.DisplayName == "" {
					for _, user := range group.Users {
						if user.ID != userID {
							names = append(names, user.Name)
						}
					}
					session.DisplayName = strings.Join(names, ",")
				}
			}
		}
	}

	return sessions, nil
}

func (s *SessionManager) Create(session *models.Session) (*models.Session, error) {
	result, err := s.getWithScope(&session.Scope)
	if err != nil {
		if err := db.Save(&session).Error; err != nil {
			return session, err
		}
	}

	log.Warn("session has been exist")
	return result, nil
}

func (s *SessionManager) Get(id string) (*models.Session, error) {
	var session models.Session
	if err := db.Where("id = ?", id).First(&session).Error; err != nil {
		return nil, err
	}

	if session.Kind == models.ScopeUser {
		user, err := models.GlobalEnvironment.Users.GetUser(session.DestinationID)
		if err != nil {
			log.Error(err)
		}
		session.DisplayName = user.Name
	}

	if session.Kind == models.ScopeGroup {
		group, err := models.GlobalEnvironment.Groups.GetGroup(session.GroupID)
		if err != nil {
			log.Error(err)
		}
		if group.DisplayName != "" {
			session.DisplayName = group.DisplayName
		} else {
			names := make([]string, 0)
			if session.DisplayName == "" {
				for _, user := range group.Users {
					names = append(names, user.Name)
				}
				session.DisplayName = strings.Join(names, ",")
			}
		}
	}

	session.Enabled = true
	if err := db.First(&models.RoleAssignment{}, &models.RoleAssignment{
		Scope: models.Scope{
			Kind:          session.Kind,
			OwnerID:       session.OwnerID,
			DestinationID: session.DestinationID,
			GroupID:       session.GroupID,
		},
	}).Error; err != nil {
		session.Enabled = false
	}

	return &session, nil
}

func (s *SessionManager) Delete(id string) error {
	return db.Delete(&models.Session{}, "id = ?", id).Error
}

func (m *SessionManager) subscribe() {
	for evt := range event.Sub("message:distribution") {
		if v, ok := evt.(*models.MessageDistributionEvent); ok {
			if v.Kind == models.ScopeUser {
				err := m.createOrUpdate(&models.Scope{
					Kind:          models.ScopeUser,
					OwnerID:       v.DestinationID,
					DestinationID: v.OwnerID,
				}, v.Content)
				if err != nil {
					log.Error(err)
				}
			} else if v.Kind == models.ScopeGroup {
				roles, err := models.GlobalEnvironment.RoleAssignments.GetRoleAssignments(
					&models.RoleAssignment{
						Scope: models.Scope{
							Kind:    models.ScopeGroup,
							GroupID: v.GroupID,
						},
					},
				)
				if err != nil {
					log.Error(err)
				}

				for _, role := range roles {
					err := m.createOrUpdate(&models.Scope{
						Kind:    models.ScopeGroup,
						OwnerID: role.OwnerID,
						GroupID: v.GroupID,
					}, v.Content)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}
}

func (m *SessionManager) createOrUpdate(scope *models.Scope, message string) error {
	session, err := m.getWithScope(scope)

	if err != nil {
		if err := db.Save(&models.Session{
			Scope: models.Scope{
				Kind:          scope.Kind,
				OwnerID:       scope.OwnerID,
				DestinationID: scope.DestinationID,
				GroupID:       scope.GroupID,
			},
		}).Error; err != nil {
			return err
		}
	}

	return db.Model(session).Updates(map[string]interface{}{
		"latest_message": message,
	}).Error

}

func (s *SessionManager) getWithScope(scope *models.Scope) (*models.Session, error) {
	var session models.Session

	if err := db.Where(&models.Session{
		Scope: models.Scope{
			Kind:          scope.Kind,
			OwnerID:       scope.OwnerID,
			DestinationID: scope.DestinationID,
			GroupID:       scope.GroupID,
		},
	}).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}
