package user

import "go-tickets/internal/user/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.Response, error) {
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	response := &dto.Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
	}
	return response, nil
}
