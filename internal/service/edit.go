package service

import "github.com/auth_test/internal/repository"

type EditService struct {
	repo repository.Edit
}

func NewEditService(repo repository.Edit) *EditService {
	return &EditService{repo: repo}
}

func (s *EditService) DeleteUser(username string) error {
	err := s.repo.DeleteUser(username)
	if err != nil {
		return err
	}
	return nil
}

func (s *EditService) UpdateUsername(usernameOld, usernameNew string) (int, error) {
	id, err := s.repo.UpdateUsername(usernameOld, usernameNew)
	if err != nil {
		return 0, err
	}

	return id, err

}
