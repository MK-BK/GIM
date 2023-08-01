package service

import (
	"GIM/models"
	"GIM/pkg/config"
)

type MomentManager struct{}

var _ models.Moments = (*MomentManager)(nil)

func NewMomentManager() *MomentManager {
	return &MomentManager{}
}

func (m *MomentManager) List(userID string) ([]*models.Moment, error) {
	moments := make([]*models.Moment, 0)
	if err := config.DB.Model(&models.Moment{}).
		Joins("left join role_assignments on role_assignments.destination_id = moments.owner_id").
		Where("role_assignments.kind = ? AND role_assignments.owner_id = ?", models.ScopeUser, userID).
		Or("moments.owner_id = ?", userID).
		Order("created_at DESC").
		Scan(&moments).Error; err != nil {
		return nil, err
	}

	for _, moment := range moments {
		user, err := models.GlobalEnvironment.Users.GetUser(moment.OwnerID)
		if err != nil {
			log.Error(err)
			continue
		}

		var comments []models.Comment
		if err := config.DB.Where("moment_id = ?", moment.ID).Find(&comments).Error; err != nil {
			log.Error(err)
		}

		moment.Comments = comments
		moment.DisplayName = user.Name
	}

	return moments, nil
}

func (m *MomentManager) Create(moment *models.Moment) error {
	return config.DB.Save(moment).Error
}

func (m *MomentManager) Get(id string) (*models.Moment, error) {
	var moment models.Moment
	if err := config.DB.Find(&moment, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &moment, nil
}

func (m *MomentManager) Delete(id string) error {
	return config.DB.Delete(&models.Moment{}, "id = ?", id).Error
}
