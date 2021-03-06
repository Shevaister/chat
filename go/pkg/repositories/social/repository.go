package social

import (
	sm "go/pkg/models/social" //sm = social model

	"gorm.io/gorm"
)

type SocialRepository struct {
	DB *gorm.DB
}

// status 1 = friend
// status 0 = blocked

func ProvideSocialRepostiory(DB *gorm.DB) SocialRepository {
	DB.AutoMigrate(&sm.Social{})
	return SocialRepository{DB: DB}
}

func (s *SocialRepository) AddUser(id1, id2 int) bool {
	var social sm.Social
	if s.checkExist(id1, id2, 1, social) {
		return false
	}
	social = sm.Social{User1ID: id1, User2ID: id2, Status: 1}
	s.DB.Create(&social)
	return true
}

func (s *SocialRepository) BlockUser(id1, id2 int) bool {
	var social sm.Social
	if s.checkExist(id1, id2, 0, social) {
		return false
	}
	social = sm.Social{User1ID: id1, User2ID: id2, Status: 0}
	s.DB.Create(&social)
	return true
}

func (s *SocialRepository) UnfriendUser(id1, id2 int) {
	var social sm.Social
	s.DB.Where("user1_id = ? AND user2_id = ? AND status = 1", id1, id2, 1).Delete(&social)
}

func (s *SocialRepository) GetFriends(id int) []sm.Social {
	var social []sm.Social
	s.DB.Where("user1_id = ? AND status = 1", id).Find(&social)
	return social
}

func (s *SocialRepository) CheckBlocked(id1, id2 int) bool {
	var social sm.Social
	return s.checkExist(id1, id2, 0, social)
}

// Utils

func (s *SocialRepository) checkExist(id1, id2, status int, social sm.Social) bool {
	err := s.DB.Where("user1_id = ? AND user2_id = ? and status = ?", id1, id2, status).Take(&social).Error
	return err == nil
}
