package repositories

import "golang-echo/models"

type Admin interface {
	GetOneIDAdmin(id int) (models.Admin, error)
	GetAllAdmin() ([]models.Admin, error)
	SaveAdmin(admin models.Admin) (models.Admin, error)
}

func (r *repository) GetOneIDAdmin(id int) (models.Admin, error) {

	var admin models.Admin

	err := r.db.First(&admin, "ID=?", id).Error

	return admin, err
}
func (r *repository) GetAllAdmin() ([]models.Admin, error) {

	var admin []models.Admin

	err := r.db.Find(&admin).Error

	return admin, err
}

func (r *repository) SaveAdmin(admin models.Admin) (models.Admin, error) {
	err := r.db.Save(&admin).Error
	return admin, err
}
