package comment

import "errors"

type Service interface {
	GetAllComment() ([]Comment, error)
	GetCommentById(id int) (Comment, error)
	CreateComment(input CreateCommentInput) (Comment, error)
	UpdateComment(input CreateCommentInput, id int) (Comment, error)
	DeleteComment(id int, userId int) (Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllComment() ([]Comment, error) {
	comments, err := s.repository.All()
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (s *service) GetCommentById(id int) (Comment, error) {
	comment, err := s.repository.FindById(id)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (s *service) CreateComment(input CreateCommentInput) (Comment, error) {
	comment := Comment{}
	comment.UserId = input.User.Id
	comment.Message = input.Message
	comment.PhotoId = input.PhotoId

	newComment, err := s.repository.Save(comment)
	if err != nil {
		return comment, err
	}

	return newComment, nil
}

func (s *service) UpdateComment(input CreateCommentInput, id int) (Comment, error) {
	comment, err := s.repository.FindById(id)
	if err != nil {
		return comment, err
	}

	if comment.Id == 0 {
		return comment, errors.New("comment not found")
	}

	if comment.UserId != input.User.Id {
		return comment, errors.New("not your comment")
	}

	comment.UserId = input.User.Id
	comment.Message = input.Message
	comment.PhotoId = input.PhotoId

	updatedComment, err := s.repository.Update(comment)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func (s *service) DeleteComment(id int, userId int) (Comment, error) {
	comment, err := s.repository.FindById(id)
	if err != nil {
		return comment, err
	}

	if comment.Id == 0 {
		return comment, errors.New("comment not found")
	}

	if comment.UserId != userId {
		return comment, errors.New("not your photo")
	}

	_, err = s.repository.Delete(comment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
