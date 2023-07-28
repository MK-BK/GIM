package user

import (
	"errors"
	"fmt"
	"time"

	"GIM/models"
	"GIM/pkg/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	redisClient *redis.Client
)

type UserManager struct{}

var _ models.Users = (*UserManager)(nil)

func NewUserManager() *UserManager {
	db = config.GetDB()
	redisClient = config.GetRedisClient()

	m := &UserManager{}

	go m.subscribe()

	return m
}

func (m *UserManager) Login(user *models.User) (string, error) {
	var token string
	if err := db.Where(&models.User{Name: user.Name, Password: user.Password}).
		First(user).Error; err != nil {
		return token, err
	}

	token, err := GenerateToken(user.Name, user.ID, 24*time.Hour)
	if err != nil {
		return token, err
	}

	redisClient.Set(user.ID, token, 10*time.Minute)

	return token, nil
}

func (m *UserManager) Logout(userID string) error {
	return redisClient.Del(userID).Err()
}

func (m *UserManager) CreateUser(user *models.User) error {
	err := db.Where(&models.User{Name: user.Name, Password: user.Password}).
		First(&models.User{}).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Email Duplicate")
	}

	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (m *UserManager) GetUser(id string) (*models.User, error) {
	var user models.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserManager) ListUsers(options *models.UserListOptions) ([]*models.User, error) {
	var users []*models.User
	if err := db.Find(&users, "name LIKE ?", fmt.Sprintf("%%%s%%", options.Search)).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserManager) UpdateUser(id string, patch *models.UserPatch) error {
	user, err := m.GetUser(id)
	if err != nil {
		return err
	}

	return db.Model(&user).Updates(map[string]interface{}{
		"name":     patch.Name,
		"password": patch.Password,
		"email":    patch.Email,
	}).Error
}

func GenerateToken(userName, userID string, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.LoginClaims{
		UserName: userName,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})

	return token.SignedString([]byte(models.Conf.SecretKey))
}
