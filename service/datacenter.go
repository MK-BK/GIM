package service

import (
	"GIM/models"
	"GIM/pkg/config"
	"GIM/pkg/event"
	"errors"

	Redis "github.com/go-redis/redis"
)

type DataCenterManager struct {
	Redis *Redis.Client
}

func NewDataCenterManager(env *models.Environment) *DataCenterManager {
	dataCenterManager := &DataCenterManager{
		// Redis: redisClient,
		// Publisher: pubsub.NewPublisher(time.Second*5, 1024),
	}

	// go dataCenterManager.InitPubsub()
	return dataCenterManager
}

// func (m *DataCenterManager) InitPubsub() {
// 	// subscribe action event
// 	pubsub := m.Redis.Subscribe("active")

// 	for msg := range pubsub.Channel() {
// 		// time.Sleep(time.Second)
// 		var redisSendSuccess bool = true

// 		for _, msg := range m.Redis.LRange(msg.Payload, 0, -1).Val() {
// 			var redisMessage models.Message
// 			if err := json.Unmarshal([]byte(msg), &redisMessage); err != nil {
// 				log.Error(err)
// 				continue
// 			}

// 			// TODO: save && distribution message

// 			// var resultRequestMessage models.Message
// 			// resultRequestMessage.DestinationID = redisMessage.OwnerID
// 			// resultRequestMessage.OwnerID = redisMessage.DestinationID
// 			// resultRequestMessage.Content = redisMessage.Content

// 			// session, err := Env.DataCenterManager.GetOrCreateSession(resultRequestMessage)
// 			// if err != nil {
// 			// 	log.Errorf("send user message error:", err)
// 			// }
// 			// if err := Env.DataCenterManager.UpdateSession(redisMessage, *session); err != nil {
// 			// 	log.Errorf("pubsub HandlerMessage Error:", err)
// 			// }
// 		}

// 		if redisSendSuccess {
// 			m.Redis.Del(msg.Payload)
// 			// log.Info("redis message send to user success")
// 		}
// 	}
// }

func (m *DataCenterManager) GetMessages(sessioID string) ([]*models.Message, error) {
	var session models.Session
	if err := config.DB.First(&session, "id = ?", sessioID).Error; err != nil {
		return nil, err
	}

	messages := make([]*models.Message, 0)
	if session.Kind == models.ScopeUser {
		err := config.DB.Where("owner_id = ? AND destination_id = ? AND created_at >= ?", session.OwnerID, session.DestinationID, session.CreatedAt).
			Or("owner_id = ? AND destination_id = ?", session.DestinationID, session.OwnerID).
			Order("created_at ASC").
			Find(&messages).Error
		if err != nil {
			return nil, err
		}
	} else if session.Kind == models.ScopeGroup {
		err := config.DB.Where("group_id = ? AND created_at >= ?", session.GroupID, session.CreatedAt).
			Order("created_at ASC").
			Find(&messages).Error
		if err != nil {
			return nil, err
		}
	}

	for _, message := range messages {
		user, err := models.GlobalEnvironment.Users.GetUser(message.OwnerID)
		if err != nil {
			log.Error(err)
			continue
		}
		message.OwnerName = user.Name
	}

	return messages, nil
}

func (m *DataCenterManager) HandlerMessage(message *models.Message) error {
	if err := checkMessageAuth(message); err != nil {
		message.Error = err.Error()
	}

	if err := config.DB.Create(message).Error; err != nil {
		return err
	}

	if message.Error == "" {
		event.Pub(&models.MessageDistributionEvent{Message: message})
	}
	return nil
}

func checkMessageAuth(message *models.Message) error {
	role := &models.RoleAssignment{
		Scope: message.Scope,
	}

	roles, err := models.GlobalEnvironment.RoleAssignments.GetRoleAssignments(role)
	if err != nil {
		return err
	}

	if len(roles) == 0 {
		return errors.New("No Role")
	}

	return nil
}
