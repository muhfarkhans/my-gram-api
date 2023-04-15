package photo

import "gorm.io/gorm"

type Repository interface {
	All() ([]Photo, error)
	Save(photo Photo) (Photo, error)
	First(id int) (Photo, error)
	Update(photo Photo) (Photo, error)
	Delete(photo Photo) (Photo, error)
	FindById(id int) (Photo, error)
	FindByUserId(userId int) ([]Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) All() ([]Photo, error) {
	var photos []Photo
	err := r.db.Preload("User").Find(&photos).Error
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (r *repository) Save(photo Photo) (Photo, error) {
	err := r.db.Create(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) First(id int) (Photo, error) {
	var photo Photo
	err := r.db.Preload("User").Where("id = ?", id).Find(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) Update(photo Photo) (Photo, error) {
	err := r.db.Save(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) Delete(photo Photo) (Photo, error) {
	err := r.db.Delete(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) FindById(id int) (Photo, error) {
	var photo Photo
	err := r.db.Preload("User").Where("id = ?", id).Find(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) FindByUserId(userId int) ([]Photo, error) {
	var photos []Photo
	err := r.db.Where("user_id = ?", userId).Preload("User").Find(&photos).Error
	if err != nil {
		return photos, err
	}

	return photos, nil
}
