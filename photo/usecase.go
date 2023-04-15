package photo

import "errors"

type Service interface {
	GetAllPhoto() ([]Photo, error)
	GetPhotoById(id int) (Photo, error)
	CreatePhoto(input CreatePhotoInput) (Photo, error)
	UpdatePhoto(input CreatePhotoInput, id int) (Photo, error)
	DeletePhoto(id int, userId int) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllPhoto() ([]Photo, error) {
	photos, err := s.repository.All()
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (s *service) GetPhotoById(id int) (Photo, error) {
	photo, err := s.repository.FindById(id)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (s *service) CreatePhoto(input CreatePhotoInput) (Photo, error) {
	photo := Photo{}
	photo.UserId = input.User.Id
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl

	newPhoto, err := s.repository.Save(photo)
	if err != nil {
		return photo, err
	}

	return newPhoto, nil
}

func (s *service) UpdatePhoto(input CreatePhotoInput, id int) (Photo, error) {
	photo, err := s.repository.FindById(id)
	if err != nil {
		return photo, err
	}

	if photo.Id == 0 {
		return photo, errors.New("photo not found")
	}

	if photo.UserId != input.User.Id {
		return photo, errors.New("not your photo")
	}

	photo.UserId = input.User.Id
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl

	updatedPhoto, err := s.repository.Update(photo)
	if err != nil {
		return updatedPhoto, err
	}

	return updatedPhoto, nil
}

func (s *service) DeletePhoto(id int, userId int) (Photo, error) {
	photo, err := s.repository.FindById(id)
	if err != nil {
		return photo, err
	}

	if photo.Id == 0 {
		return photo, errors.New("photo not found")
	}

	if photo.UserId != userId {
		return photo, errors.New("not your photo")
	}

	_, err = s.repository.Delete(photo)
	if err != nil {
		return photo, err
	}

	return photo, nil
}
