package social

import "gorm.io/gorm"

type SocialRepository struct {
	DB *gorm.DB
}

func ProvideSocialRepostiory(DB *gorm.DB) SocialRepository {
	DB.AutoMigrate(&Social{})
	return SocialRepository{DB: DB}
}

func (s *SocialRepository) AddUser(id1, id2 int) bool {
	var social Social
	if !checkExist(s, id1, id2, 1, social) {
		return false
	}
	social = Social{User1ID: id1, User2ID: id2, Status: 1}
	s.DB.Create(&social)
	return true
}

func (s *SocialRepository) BlockUser(id1, id2 int) bool {
	var social Social
	if !checkExist(s, id1, id2, 0, social) {
		return false
	}
	social = Social{User1ID: id1, User2ID: id2, Status: 0}
	s.DB.Create(&social)
	return true
}

func (s *SocialRepository) UnfriendUser(id1, id2 int) {
	var social Social
	s.DB.Where("user1_id = ? AND user2_id = ? AND status = 1", id1, id2, 1).Delete(&social)
}

func (s *SocialRepository) GetFriends(id int) []Social {
	var social []Social
	s.DB.Where("user1_id = ? AND status = 1", id).Find(&social)
	return social
}

// Utils

func checkExist(s *SocialRepository, id1, id2, status int, social Social) bool {
	if s.DB.Where("user1_id = ? AND user2_id = ? and status = ?", id1, id2, status).Take(&social).Error == nil {
		return false
	}
	return true
}
