package socialmedia

import "errors"

type Service interface {
	GetAllSocialmedia() ([]Socialmedia, error)
	GetSocialmediaById(id int) (Socialmedia, error)
	CreateSocialmedia(input CreateSocialMediaInput) (Socialmedia, error)
	UpdateSocialmedia(input CreateSocialMediaInput, id int) (Socialmedia, error)
	DeleteSocialmedia(id int, userId int) (Socialmedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllSocialmedia() ([]Socialmedia, error) {
	socialmedias, err := s.repository.All()
	if err != nil {
		return socialmedias, err
	}

	return socialmedias, nil
}

func (s *service) GetSocialmediaById(id int) (Socialmedia, error) {
	socialmedia, err := s.repository.FindById(id)
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (s *service) CreateSocialmedia(input CreateSocialMediaInput) (Socialmedia, error) {
	socialmedia := Socialmedia{}
	socialmedia.UserId = input.User.Id
	socialmedia.Name = input.Name
	socialmedia.SocialMediaUrl = input.SocialMediaUrl

	newSocialmedia, err := s.repository.Save(socialmedia)
	if err != nil {
		return socialmedia, err
	}

	return newSocialmedia, nil
}

func (s *service) UpdateSocialmedia(input CreateSocialMediaInput, id int) (Socialmedia, error) {
	socialmedia, err := s.repository.FindById(id)
	if err != nil {
		return socialmedia, err
	}

	if socialmedia.Id == 0 {
		return socialmedia, errors.New("socialmedia not found")
	}

	if socialmedia.UserId != input.User.Id {
		return socialmedia, errors.New("not your socialmedia")
	}

	socialmedia.UserId = input.User.Id
	socialmedia.Name = input.Name
	socialmedia.SocialMediaUrl = input.SocialMediaUrl

	updatedSocialmedia, err := s.repository.Update(socialmedia)
	if err != nil {
		return updatedSocialmedia, err
	}

	return updatedSocialmedia, nil
}

func (s *service) DeleteSocialmedia(id int, userId int) (Socialmedia, error) {
	socialmedia, err := s.repository.FindById(id)
	if err != nil {
		return socialmedia, err
	}

	if socialmedia.Id == 0 {
		return socialmedia, errors.New("socialmedia not found")
	}

	if socialmedia.UserId != userId {
		return socialmedia, errors.New("not your socialmedia")
	}

	_, err = s.repository.Delete(socialmedia)
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}
