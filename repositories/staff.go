package repositories

import "golang-echo/models"




//query

type Staff interface {
	GetOneIDStaff(id int) (models.Staff, error)
	GetAllStaff() ([]models.Staff, error)
	SaveStaff(staff models.Staff) (models.Staff, error)
}

func (r *repository) GetOneIDStaff(id int) (models.Staff, error) {

	var staff models.Staff

	err := r.db.First(&staff, "ID=?", id).Error

	return staff, err
}
func (r *repository) GetAllStaff() ([]models.Staff, error) {

	var staff []models.Staff

	err := r.db.Find(&staff).Error

	return staff, err
}

func (r *repository) SaveStaff(staff models.Staff) (models.Staff, error) {
	err := r.db.Save(&staff).Error
	return staff, err
}
