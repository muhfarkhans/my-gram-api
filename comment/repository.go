package comment

import "gorm.io/gorm"

type Repository interface {
	All() ([]Comment, error)
	Save(comment Comment) (Comment, error)
	First(id int) (Comment, error)
	Update(comment Comment) (Comment, error)
	Delete(comment Comment) (Comment, error)
	FindById(id int) (Comment, error)
	FindByUserId(userId int) ([]Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) All() ([]Comment, error) {
	var comments []Comment
	err := r.db.Preload("User").Preload("Photo").Preload("Photo.User").Find(&comments).Error
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (r *repository) Save(comment Comment) (Comment, error) {
	err := r.db.Create(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) First(id int) (Comment, error) {
	var comment Comment
	err := r.db.Where("id = ?", id).Find(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) Update(comment Comment) (Comment, error) {
	err := r.db.Save(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) Delete(comment Comment) (Comment, error) {
	err := r.db.Delete(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) FindById(id int) (Comment, error) {
	var comment Comment
	err := r.db.Preload("User").Preload("Photo").Where("id = ?", id).Find(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) FindByUserId(userId int) ([]Comment, error) {
	var comments []Comment
	err := r.db.Where("user_id = ?", userId).Preload("User").Find(&comments).Error
	if err != nil {
		return comments, err
	}

	return comments, nil
}
