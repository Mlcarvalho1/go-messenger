package services

import (
	"go.messenger/database"
	"go.messenger/models"
)

type CreateGroupRequest struct {
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description"`
	Members     []uint      `json:"members"`
}

func CreateGroup(req CreateGroupRequest) (*models.Group, error) {
	group := &models.Group{
		Name:        req.Name,
		Description: req.Description,
	}

	tx := database.DB.Db.Begin()

	if err := tx.Create(group).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, memberID := range req.Members {
		groupMember := &models.GroupMember{
			GroupID: group.ID,
			UserID:  memberID,
		}
		if err := tx.Create(groupMember).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := database.DB.Db.Preload("GroupMembers").Preload("GroupMembers.User").First(group, group.ID).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func GetGroupChatsByUserID(userID int) ([]models.Group, error) {
    var groups []models.Group

    // Subconsulta para encontrar os grupos em que o usuário é membro
    subQuery := database.DB.Db.Model(&models.GroupMember{}).Select("group_id").Where("user_id = ?", userID)
	if(subQuery.Error != nil) {
		return nil, subQuery.Error
	}

    // Buscar os grupos usando a subconsulta
    result := database.DB.Db.Where("id IN (?)", subQuery).Preload("GroupMembers").Preload("GroupMembers.User").Find(&groups)
    if result.Error != nil {
        return nil, result.Error
    }

    return groups, nil
}