package user

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *Service {
	return &Service{db}
}

func (s *Service) CreateUser() string {
	user := User{Name: "John"}
	s.db.Create(&user)
	return "User created"
}

func (s *Service) GetUserById(id string) (*User, error) {
	var user User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
