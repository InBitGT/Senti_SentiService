package address

import "gorm.io/gorm"

type Repository interface {
	Create(a *Address) error
	GetByID(id uint) (*Address, error)
	GetAll() ([]Address, error)
	Update(id uint, a *Address) error
	Delete(id uint) error
	HardDelete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(a *Address) error {
	return r.db.Create(a).Error
}

func (r *repository) GetByID(id uint) (*Address, error) {
	var a Address
	if err := r.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *repository) GetAll() ([]Address, error) {
	var list []Address
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *repository) Update(id uint, a *Address) error {
	return r.db.Model(&Address{}).Where("id_address = ?", id).Updates(a).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Address{}, id).Error
}

func (r *repository) HardDelete(id uint) error {
	return r.db.Unscoped().Where("id_address = ?", id).Delete(&Address{}).Error
}
