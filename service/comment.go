package service

import (
	"GIM/models"
	"GIM/pkg/config"
)

type commentManager struct{}

var _ models.Comments = (*commentManager)(nil)

func NewCommonManager() *commentManager {
	return &commentManager{}
}

func (m *commentManager) List(id string) ([]*models.Comment, error) {
	var comments []*models.Comment
	if err := config.DB.Where("id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (m *commentManager) Create(comment *models.Comment) error {
	return config.DB.Create(comment).Error
}

func (m *commentManager) Delete(id string) error {
	return config.DB.Delete(&models.Comment{}, "id = ?", id).Error
}
