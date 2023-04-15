package socialmedia

import "gorm.io/gorm"

type Repository interface {
	All() ([]Socialmedia, error)
	Save(socialmedia Socialmedia) (Socialmedia, error)
	First(id int) (Socialmedia, error)
	Update(socialmedia Socialmedia) (Socialmedia, error)
	Delete(socialmedia Socialmedia) (Socialmedia, error)
	FindById(id int) (Socialmedia, error)
	FindByUserId(userId int) ([]Socialmedia, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) All() ([]Socialmedia, error) {
	var socialmedias []Socialmedia
	err := r.db.Preload("User").Find(&socialmedias).Error
	if err != nil {
		return socialmedias, err
	}

	return socialmedias, nil
}

func (r *repository) Save(socialmedia Socialmedia) (Socialmedia, error) {
	err := r.db.Create(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *repository) First(id int) (Socialmedia, error) {
	var socialmedia Socialmedia
	err := r.db.Where("id = ?", id).Find(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *repository) Update(socialmedia Socialmedia) (Socialmedia, error) {
	err := r.db.Save(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *repository) Delete(socialmedia Socialmedia) (Socialmedia, error) {
	err := r.db.Delete(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *repository) FindById(id int) (Socialmedia, error) {
	var socialmedia Socialmedia
	err := r.db.Preload("User").Where("id = ?", id).Find(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *repository) FindByUserId(userId int) ([]Socialmedia, error) {
	var socialmedias []Socialmedia
	err := r.db.Where("user_id = ?", userId).Preload("User").Find(&socialmedias).Error
	if err != nil {
		return socialmedias, err
	}

	return socialmedias, nil
}
