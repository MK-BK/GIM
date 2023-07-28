package service

import (
	"GIM/models"
)

type commentManager struct{}

var _ models.Comments = (*commentManager)(nil)

func NewCommonManager() *commentManager {
	return &commentManager{}
}

func (m *commentManager) List(id string) ([]*models.Comment, error) {
	var comments []*models.Comment
	if err := db.Where("id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (m *commentManager) Create(comment *models.Comment) error {
	return db.Create(comment).Error
}

func (m *commentManager) Delete(id string) error {
	return db.Delete(&models.Comment{}, "id = ?", id).Error
}
