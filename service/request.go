package service

import (
	"GIM/models"
	"GIM/pkg/event"
)

type RequestManager struct{}

var _ models.Requests = (*RequestManager)(nil)

func NewRequestManager() *RequestManager {
	return &RequestManager{}
}

func (m *RequestManager) CreateRequest(request *models.Request) error {
	return db.Save(&request).Error
}

func (m *RequestManager) GetRequest(id string) (*models.Request, error) {
	var request models.Request
	if err := db.Where("id = ?", id).Find(&request).Error; err != nil {
		return nil, err
	}

	return &request, nil
}

func (m *RequestManager) ListRequests(userID string) ([]*models.Request, error) {
	var requests []*models.Request
	if err := db.Where("destination_id = ?", userID).
		Order("created_at DESC").
		Find(&requests).Error; err != nil {
		return nil, err
	}

	for _, request := range requests {
		user, err := models.GlobalEnvironment.Users.GetUser(request.OwnerID)
		if err != nil {
			log.Error(err)
			continue
		}
		request.OwnerName = user.Name
	}

	return requests, nil
}

func (m *RequestManager) UpdateRequest(userID, id string) error {
	var request models.Request
	if err := db.Where("id = ?", id).First(&request).Error; err != nil {
		return err
	}

	// if err := m.Verify(&request, userID); err != nil {
	// 	return err
	// }

	if err := db.Model(&request).Update("status", true).Error; err != nil {
		return err
	}

	log.Info("publish ack request event to module of users")

	event.Pub(&models.RequestUpdateEvent{
		Request: &request,
	})
	return nil
}

// TODO: use const var replace magic number -- 0: user, 1: room
func (m *RequestManager) Verify(request *models.Request, userID string) error {
	// switch request.Kind {
	// case 0:
	// 	if userID != request.OwnerID {
	// 		return fmt.Errorf("%s not permission to resolve from %s's request", userID, request.DestinationID)
	// 	}
	// case 1:
	// 	room, err := m.GroupInterface.GetGroup(request.RoomID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if userID != room.ManagerID {
	// 		return fmt.Errorf("%s not the manager of %s room", userID, room.DisplayName)
	// 	}
	// }

	return nil
}

func (m *RequestManager) DeleteRequest(id string) error {
	return db.Delete(&models.Request{}, "id = ?", id).Error
}

func (*RequestManager) GetDescription(request *models.Request) string {
	var description string
	// if request.Kind == models.USERMESSAGE {
	// 	user, err := models.GlobalEnvironment.UserInterface.GetUser(request.OwnerID)
	// 	if err != nil {
	// 		log.Error(err)
	// 	}

	// 	description = fmt.Sprintf("%s 请求添加好友", user.Name)
	// }

	// if request.Kind == models.GROUPMESSAGE {
	// 	visitUser, err := models.GlobalEnvironment.UserInterface.GetUser(request.InviteID)
	// 	if err != nil {
	// 		log.Error(err)
	// 	}

	// 	hasVistUser, err := models.GlobalEnvironment.UserInterface.GetUser(request.OwnerID)
	// 	if err != nil {
	// 		log.Error(err)
	// 	}

	// 	description = fmt.Sprintf("%s 邀请 %s 进入群聊", visitUser.Name, hasVistUser.Name)

	// }
	return description
}
